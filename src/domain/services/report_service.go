package services

import (
	"quake_report/src/domain/models"
	"quake_report/src/domain/ports"
)

type reportService struct{}

func NewReportService() ports.ReportService {
	return &reportService{}
}

func (s reportService) Generate(games map[string]models.Game) models.Report {
	var rankingPlayers = make(map[string]int)
	var causesOfDeath = make(map[string]int)
	for _, game := range games {
		rankingPlayers = s.consolidate(rankingPlayers, game.Kills)
		causesOfDeath = s.consolidate(causesOfDeath, game.CausesOfDeath)
	}
	report := models.Report{}
	report.Games = games
	report.RankingCausesOfDeath = causesOfDeath
	report.RankingPlayers = rankingPlayers
	return report
}

func (s reportService) consolidate(ranking map[string]int, source map[string]int) map[string]int {
	for key, score := range source {
		ranking[key] += score
	}
	return ranking
}
