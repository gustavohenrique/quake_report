package collections

import (
	"fmt"
	"sort"
)

func ConvertMapSortedSlice(unsorted map[string]int) []string {
	type Ranking struct {
		Key   string
		Score int
	}
	rankings := make([]Ranking, 0, len(unsorted))
	for key, score := range unsorted {
		rankings = append(rankings, Ranking{Key: key, Score: score})
	}
	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].Score > rankings[j].Score
	})
	var sorted []string
	for _, ranking := range rankings {
		sorted = append(sorted, fmt.Sprintf("%s: %d", ranking.Key, ranking.Score))
	}
	return sorted
}
