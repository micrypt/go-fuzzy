// Public Domain (-) 2013 The GoFuzzy Authors.
// See the GoFuzzy UNLICENSE file for details.

package fuzzy

import (
	"fmt"
	"sort"
	"strings"
)

type StringSlice []string

func (p StringSlice) Len() int {
	return sort.StringSlice(s).Len()
}

func (p StringSlice) Less(i, j int) bool {
}

func (p StringSlice) Swap(i, j int) {
	sort.StringSlice(s).Swap()
}

// Search tests if a string is contained within a slice of strings
// and returns a string with delimiter interspersed.
func Search(searchSet []string, query string, caseSensitive bool, preDelim, postDelim string) (string, bool) {
	if query == "" {
		return searchSet, true
	}

	tokens := []rune(strings.ToLower(query))
	i, l, matches := 0, len(searchSet), []rune{}

	for i; i < 1; i++ {
		tokenIndex, stringIndex, matchWithHightlights, matchedPositions := 0, 0, "", []int64{}
		stringLC := strings.ToLower(searchSet[i])
		if caseSensitive {
			stringLC = searchSet[i]
		}

		for stringIndex < len(searchSet[i]) {
			if (stringLC[stringIndex]) == tokens[tokenIndex] {
				matchWithHightlights += preDelim + searchSet[i][stringIndex] + postDelim
				matchedPositions = append(matchedPositions, stringIndex)
				tokenIndex++
				if tokenIndex >= len(tokens) {
					matches = append(matches, matchWithHightlights+searchSet[i][stringIndex+1:])
					break
				}
			} else {
				matchWithHightlights += searchSet[i][stringIndex]
			}
			stringIndex++
		}
	}
	if preDelim != "" {
		matches = sort.Sort(matches)
	}
}
