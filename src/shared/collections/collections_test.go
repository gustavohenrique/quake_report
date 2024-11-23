package collections_test

import (
	"testing"

	"quake_report/src/shared/collections"
)

func TestConvertMapToSortedSlice(t *testing.T) {
	unsorted := map[string]int{
		"A": 4,
		"B": 8,
		"C": 0,
	}
	sorted := collections.ConvertMapSortedSlice(unsorted)
	if sorted[0] != "B: 8" {
		t.Fatalf("want B: 8, got %s", sorted[0])
	}
}
