package services_test

import (
	"maps"
	"slices"
	"sort"
	"testing"

	"quake_report/src/domain/services"
	"quake_report/src/shared/testify"
	"quake_report/src/shared/testify/assert"
)

func TestGameService(ts *testing.T) {
	const logContent = `
  0:00 ------------------------------------------------------------
  0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0
 15:00 Exit: Timelimit hit.
 20:34 ClientConnect: 2
 20:34 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\xian/default\hmodel\xian/default\g_redteam\\g_blueteam\\c1\4\c2\5\hc\100\w\0\l\0\tt\0\tl\0
 20:37 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\uriel/zael\hmodel\uriel/zael\g_redteam\\g_blueteam\\c1\5\c2\5\hc\100\w\0\l\0\tt\0\tl\0
 20:37 ClientBegin: 2
 20:37 ShutdownGame:
 20:37 ------------------------------------------------------------
 20:37 ------------------------------------------------------------
 20:37 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\bot_minplayers\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0
 20:38 ClientConnect: 2
 20:38 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\uriel/zael\hmodel\uriel/zael\g_redteam\\g_blueteam\\c1\5\c2\5\hc\100\w\0\l\0\tt\0\tl\0
 20:38 ClientBegin: 2
 20:40 Item: 2 weapon_rocketlauncher
 20:40 Item: 2 ammo_rockets
 20:42 Item: 2 item_armor_body
 20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
 20:59 Item: 2 weapon_rocketlauncher
 21:04 Item: 2 ammo_shells
 21:07 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
 22:06 Kill: 2 3 7: Isgalamido killed Mocinha by MOD_ROCKET_SPLASH
 22:11 Item: 2 item_quad
 22:11 ClientDisconnect: 3
 22:18 Kill: 2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH
  1:08 Kill: 2 2 8: Isgalamido killed Mocinha by MOD_ROCKET`

	testify.It(ts, "Parse", func(tt *testing.T) {
		testify.It(tt, "Should process InitGame event return 2 games", func(t *testing.T) {
			service := services.NewGameService()
			games := service.Parse([]byte(logContent))
			assert.Equal(t, 2, len(games))
		})

		testify.It(tt, "Should process Kill event and get who is killed and causes of death", func(t *testing.T) {
			service := services.NewGameService()
			games := service.Parse([]byte(logContent))
			game1 := games["game_1"]
			assert.Equal(t, 0, game1.TotalKills)
			assert.Equal(t, 0, len(game1.Kills))
			assert.Equal(t, 0, len(game1.CausesOfDeath))

			game2 := games["game_2"]
			assert.Equal(t, 5, game2.TotalKills)
			kills := slices.Collect(maps.Keys(game2.Kills))
			assert.Equal(t, 1, len(kills))
			assert.Equal(t, "Isgalamido", kills[0])
			causes := slices.Collect(maps.Keys(game2.CausesOfDeath))
			sort.Strings(causes)
			assert.Equal(t, 3, len(causes))
			assert.Equal(t, "MOD_ROCKET", causes[0])
			assert.Equal(t, "MOD_ROCKET_SPLASH", causes[1])
			assert.Equal(t, "MOD_TRIGGER_HURT", causes[2])
		})

		testify.It(tt, "Should get all players of each game", func(t *testing.T) {
			service := services.NewGameService()
			games := service.Parse([]byte(logContent))
			game1 := games["game_1"]
			assert.Equal(t, 2, len(game1.Players))
			assert.Equal(t, "Isgalamido", game1.Players[0])
			assert.Equal(t, "Mocinha", game1.Players[1])

			game2 := games["game_2"]
			players := game2.Players
			assert.Equal(t, 2, len(players))
			assert.Equal(t, "Isgalamido", players[0])
			assert.Equal(t, "Mocinha", players[1])
		})
	})
}
