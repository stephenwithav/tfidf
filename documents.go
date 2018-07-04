package tfidf

import (
	"io"
	"sort"
)

func DocumentFrequency(docs <-chan io.Reader, docNames <-chan string) map[string][]string {
	// docCount := 0
	// docFrequencies := make(map[string]float64)
	termWheres := make(map[string][]termRelevance)

	// Keep going until the docs channel is closed.
	for doc := range docs {
		termFrequencies := TermFrequency(doc)
		docName := <-docNames

		// A map underlies termFrequencies, so we can range over it.
		// Append the doc source, which could be a url or a filename,
		// to record its location.
		for term, pct := range termFrequencies {
			// This term composes {pct}% of docName.
			termWheres[term] = append(termWheres[term], termRelevance{docName, pct})
		}
	}

	// Now termWheres contains an index of word->docList.
	// Sort documents containing each term in decreasing order of frequency.
	// And can we dispense with the termRelevance* sleight of hand?
	// I believe so.
	termLocations := make(map[string][]string)
	for term, wheres := range termWheres {
		sort.Sort(termRelevances(wheres))

		for _, where := range wheres {
			termLocations[term] = append(termLocations[term], where.docName)
		}
	}

	// We have our sorting, no inversion necessary due to the
	// inversion of Less in termRelevances.
	return termLocations
}

// Helper struct for DocumentFrequency.
type termRelevance struct {
	docName         string
	termComposition float64
}

// Simple sort.Interface implementation for termRelevances.
// Sorts in decreasing order.
type termRelevances []termRelevance

func (trs termRelevances) Len() int {
	return len(trs)
}

func (trs termRelevances) Less(i, j int) bool {
	// Sort in decreasing order.
	return trs[i].termComposition > trs[j].termComposition
}

func (trs termRelevances) Swap(i, j int) {
	trs[i], trs[j] = trs[j], trs[i]
}
