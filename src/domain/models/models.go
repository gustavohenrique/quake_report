package models

type Game struct {
	TotalKills    int
	Players       []string
	Kills         map[string]int
	CausesOfDeath map[string]int
}

type Report struct {
	Games                map[string]Game
	RankingPlayers       map[string]int
	RankingCausesOfDeath map[string]int
}
