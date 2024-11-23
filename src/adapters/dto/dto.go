package dto

type GamePresentation struct {
	TotalKills    int            `json:"total_kills"`
	Players       []string       `json:"players"`
	Kills         map[string]int `json:"kills"`
	CausesOfDeath map[string]int `json:"causes_of_death"`
}

type ReportPresentation struct {
	Games                map[string]GamePresentation `json:"games"`
	RankingPlayers       []string                    `json:"ranking_players"`
	RankingCausesOfDeath []string                    `json:"ranking_causes_of_death"`
}
