package converters

import (
	"quake_report/src/adapters/dto"
	"quake_report/src/domain/models"
	"quake_report/src/shared/collections"
)

type ReportConverter struct{}

func NewReportConverter() ReportConverter {
	return ReportConverter{}
}

func (c ReportConverter) FromModelToPresentation(input models.Report) dto.ReportPresentation {
	output := dto.ReportPresentation{}
	output.RankingCausesOfDeath = collections.ConvertMapSortedSlice(input.RankingCausesOfDeath)
	output.RankingPlayers = collections.ConvertMapSortedSlice(input.RankingPlayers)

	output.Games = make(map[string]dto.GamePresentation)
	for k, v := range input.Games {
		output.Games[k] = dto.GamePresentation{
			TotalKills:    v.TotalKills,
			Players:       v.Players,
			Kills:         v.Kills,
			CausesOfDeath: v.CausesOfDeath,
		}
	}
	return output
}
