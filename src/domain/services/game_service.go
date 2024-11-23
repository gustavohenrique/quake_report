package services

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"sort"

	"quake_report/src/domain/models"
	"quake_report/src/domain/ports"
)

const (
	initGameEvent    = "initGame"
	gameKillEvent    = "kill"
	unknownGameEvent = "unknown"
	worldPlayer      = "<world>"
	gameNamePrefix   = "game_"
)

var (
	initGameRegex = regexp.MustCompile(`InitGame:`)
	killRegex     = regexp.MustCompile(`Kill: \d+ \d+ \d+: (.+) killed (.+) by (.+)`)
)

type (
	processGameEventFunc = func(string, models.Game) models.Game
	gameService          struct {
		currentGameName string
		players         map[string]struct{}
		games           map[string]models.Game
	}
)

func NewGameService() ports.GameService {
	return &gameService{
		games:   make(map[string]models.Game),
		players: make(map[string]struct{}),
	}
}

func (s *gameService) Parse(content []byte) map[string]models.Game {
	var game = models.Game{
		Kills:         make(map[string]int),
		CausesOfDeath: make(map[string]int),
	}
	var processors = map[string]processGameEventFunc{
		initGameEvent: s.processInitGameEvent,
		gameKillEvent: s.processGameKillEvent,
		unknownGameEvent: func(string, models.Game) models.Game {
			return game
		},
	}

	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		eventType := s.getEventTypeBy(line)
		processGameEventFunc := processors[eventType]
		game = processGameEventFunc(line, game)
	}

	if s.currentGameName != "" {
		s.games[s.currentGameName] = game
	}
	games := s.addPlayersToEachGame(s.games, s.players)
	return games
}

func (s *gameService) processInitGameEvent(line string, game models.Game) models.Game {
	if s.currentGameName != "" {
		s.games[s.currentGameName] = game
	}
	s.currentGameName = fmt.Sprintf("%s%d", gameNamePrefix, len(s.games)+1)
	s.players = make(map[string]struct{})
	return models.Game{
		Kills:         make(map[string]int),
		CausesOfDeath: make(map[string]int),
	}
}

func (s *gameService) processGameKillEvent(line string, game models.Game) models.Game {
	matches := killRegex.FindStringSubmatch(line)
	killer := matches[1]
	victim := matches[2]
	cause := matches[3]

	if killer != worldPlayer {
		game.Kills[killer]++
		s.players[killer] = struct{}{}
	} else {
		game.Kills[victim]--
	}
	s.players[victim] = struct{}{}
	game.CausesOfDeath[cause]++
	game.TotalKills++
	return game
}

func (s *gameService) addPlayersToEachGame(games map[string]models.Game, players map[string]struct{}) map[string]models.Game {
	for gameName, game := range games {
		playerSet := make([]string, 0, len(players))
		for player := range players {
			if player != worldPlayer {
				playerSet = append(playerSet, player)
			}
		}
		sort.Strings(playerSet)
		game.Players = playerSet
		games[gameName] = game
	}
	return games
}

func (s *gameService) getEventTypeBy(line string) string {
	var isInitGameEvent = initGameRegex.MatchString(line)
	var isGameKillEvent = killRegex.MatchString(line)
	if isInitGameEvent {
		return initGameEvent
	}
	if isGameKillEvent {
		return gameKillEvent
	}
	return unknownGameEvent
}
