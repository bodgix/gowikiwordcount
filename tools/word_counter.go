//Package tools provides primitives for counting and sorting
//words
package tools

import (
	"sort"
	"strings"
)

// WordWithCount - word and it's count tuple
type WordWithCount struct {
	Word  string
	Count int
}

// ByCountReverse - sorts by count in reverse order
type ByCountReverse []WordWithCount

// Len - returns the length of the ByCountReverse
func (wc ByCountReverse) Len() int {
	return len(wc)
}

// Less - for comparing
func (wc ByCountReverse) Less(i, j int) bool {
	return wc[i].Count > wc[j].Count
}

// Swap - swap the two
func (wc ByCountReverse) Swap(i, j int) {
	wc[i], wc[j] = wc[j], wc[i]
}

// wordCounter - counts the occurances of words in a
// words slice.
// Returns a slice of WordWithCount ordered by count
func wordCounter(words []string) map[string]int {
	wordsCount := make(map[string]int)
	for _, word := range words {
		wordsCount[strings.ToLower(word)]++
	}
	return wordsCount
}

// TopN - returns top N most common words
func TopN(words []string, n int) []WordWithCount {
	wordsWithCnt := mapToSlice(wordCounter(words))
	sort.Sort(ByCountReverse(wordsWithCnt))
	return mergeTopN(wordsWithCnt, n)
}

func mapToSlice(wordsCount map[string]int) []WordWithCount {
	var resSlice []WordWithCount
	for word, count := range wordsCount {
		resSlice = append(resSlice, WordWithCount{word, count})
	}
	return resSlice
}
