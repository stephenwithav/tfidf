package tfidf

import (
	"io"
	"strings"
	"testing"
)

func TestDocumentFrequency(t *testing.T) {
	given := []string{
		"easter greeting: Happy easter, Easter Bunny",
		"first president: President George Washington was the first American President",
		"43rd president: George W. Bush was the 43rd American President.  He presided over 9/11, Afghanistan, Iraq, and the fall of Saddam Hussein.",
	}

	expect := map[string][]string{
		"easter":    []string{"easter greeting"},
		"happy":     []string{"easter greeting"},
		"bunny":     []string{"easter greeting"},
		"president": []string{"first president", "43rd president"},
		"american":  []string{"first president", "43rd president"},
		"iraq":      []string{"43rd president"},
	}

	docCh := make(chan io.Reader, 3)
	docNames := make(chan string, 3)

	for _, statements := range given {
		nameDoc := strings.SplitN(statements, ": ", 2)
		docCh <- strings.NewReader(nameDoc[1])
		docNames <- nameDoc[0]
	}
	close(docCh)
	close(docNames)

	termOrderings := DocumentFrequency(docCh, docNames)

	for key, values := range expect {
		got := termOrderings[key]

		if len(values) != len(got) {
			t.Errorf("got %#v, expected %#v", got, values)
		}

		for i, value := range values {
			if value != got[i] {
				t.Errorf("got %#v, expected %#v", got, values)
			}
		}
	}
}
