package services_test

import (
	"testing"

	"quake_report/src/domain/models"
	"quake_report/src/domain/services"
	"quake_report/src/shared/testify"
	"quake_report/src/shared/testify/assert"
)

func TestReportService(tt *testing.T) {
	game1 := models.Game{
		TotalKills: 2,
		Players:    []string{"Isgalamido", "Mocinha"},
		Kills: map[string]int{
			"Isgalamido": 2,
			"Mocinha":    0,
		},
		CausesOfDeath: map[string]int{
			"MOD_ROCKET": 2,
		},
	}
	game2 := models.Game{
		TotalKills: 5,
		Players:    []string{"Isgalamido", "Mocinha"},
		Kills: map[string]int{
			"Isgalamido": 3,
			"Mocinha":    2,
		},
		CausesOfDeath: map[string]int{
			"MOD_ROCKET":        2,
			"MOD_TRIGGER_HURT":  2,
			"MOD_ROCKET_SPLASH": 1,
		},
	}
	games := map[string]models.Game{"game_1": game1, "game_2": game2}

	testify.It(tt, "Should generate a consolidated report", func(t *testing.T) {
		service := services.NewReportService()
		report := service.Generate(games)
		assert.DeepEqual(t, report.Games, games)
		assert.DeepEqual(t, report.RankingCausesOfDeath, map[string]int{
			"MOD_ROCKET":        4,
			"MOD_TRIGGER_HURT":  2,
			"MOD_ROCKET_SPLASH": 1,
		})
		assert.DeepEqual(t, report.RankingPlayers, map[string]int{
			"Isgalamido": 5,
			"Mocinha":    2,
		})
	})
}
