package tfidf

import (
	"strings"
	"testing"
)

func TestTermCount(t *testing.T) {
	given := []string{
		"Happy easter, Easter Bunny",
		"President George Washington was the first American President",
	}

	expected := []map[string]float64{
		map[string]float64{"happy": 1 / 4.0, "easter": 2 / 4.0, "bunny": 1 / 4.0},
		map[string]float64{"president": 2 / 8.0, "american": 1 / 8.0, "george": 1 / 8.0, "washington": 1 / 8.0, "was": 1 / 8.0, "the": 1 / 8.0, "first": 1 / 8.0},
	}

	for i, s := range given {
		found := TermFrequency(strings.NewReader(s))

		for term, freq := range found {
			if freq != expected[i][term] {
				t.Errorf("got %f, expected %f\t[%s]", freq, expected[i][term], term)
			}
		}
	}
}
