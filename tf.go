package tfidf

import (
	"bufio"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

// scanAlphaNum is a split function for a Scanner that returns each
// alphanumeric token of text, with surrounding spaces deleted. It will
// never return an empty string.
func scanAlphaNum(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip invalid leading characters.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if isValid(r) {
			break
		}

	}

	// Scan until invalid character.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if !isValid(r) {
			return i + width, data[start:i], nil
		}

	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	// Request more data.
	return start, nil, nil
}

// isValid is a function that tells whether a rune is alphanumeric.
func isValid(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

type TermFrequencies map[string]float64

// TermFrequency is a function that returns the term-frequency of the
// given io.Reader.  No stop-words are eliminated.
func TermFrequency(doc io.Reader) TermFrequencies {
	scanner := bufio.NewScanner(doc)
	scanner.Split(scanAlphaNum)

	// We start with no words.
	termCount := 0
	terms := make(map[string]int)

	// Count the term totals.  TODO: Add an if-block.  if tok, invalid := invalidTokens[lowerCasedText]; !invalid { code here }
	for scanner.Scan() {
		terms[strings.ToLower(scanner.Text())]++
		termCount++
	}

	// Now calculate their frequency as a percentage.
	tFrequency := make(map[string]float64, len(terms))
	for term, count := range terms {
		tFrequency[term] = float64(count) / float64(termCount)
	}

	return TermFrequencies(tFrequency)
}
