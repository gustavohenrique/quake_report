package converters_test

import (
	"quake_report/src/adapters/converters"
	"quake_report/src/domain/models"
	"quake_report/src/shared/testify/assert"
	"testing"
)

func TestReportConverter(t *testing.T) {
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
	input := models.Report{
		Games: games,
		RankingPlayers: map[string]int{
			"Isgalamido": 5,
			"Mocinha":    2,
		},
		RankingCausesOfDeath: map[string]int{
			"MOD_ROCKET":        4,
			"MOD_TRIGGER_HURT":  2,
			"MOD_ROCKET_SPLASH": 1,
		},
	}
	output := converters.NewReportConverter().FromModelToPresentation(input)
	assert.DeepEqual(t, output.RankingCausesOfDeath, []string{"MOD_ROCKET: 4", "MOD_TRIGGER_HURT: 2", "MOD_ROCKET_SPLASH: 1"})
	assert.DeepEqual(t, output.RankingPlayers, []string{"Isgalamido: 5", "Mocinha: 2"})
}
