package ports

import "quake_report/src/domain/models"

type GameService interface {
	Parse(content []byte) map[string]models.Game
}

type ReportService interface {
	Generate(games map[string]models.Game) models.Report
}
