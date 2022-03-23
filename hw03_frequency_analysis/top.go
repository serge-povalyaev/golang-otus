package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Stat struct {
	Word  string
	Count int
}

func Top10(str string) []string {
	counts := GetSortingStat(str, 10)
	result := make([]string, 0, len(counts))
	for _, wordStat := range counts {
		result = append(result, wordStat.Word)
	}

	return result
}

func GetSortingStat(str string, resultMaxLength int) []Stat {
	words := StringExplode(str)
	wordsCounts := CalculateCounts(words)

	sort.Slice(wordsCounts, func(i, j int) bool {
		return Sort(wordsCounts[i], wordsCounts[j])
	})

	if len(wordsCounts) <= resultMaxLength {
		return wordsCounts
	}

	return wordsCounts[0:resultMaxLength]
}

func StringExplode(str string) []string {
	return strings.Fields(str)
}

func CalculateCounts(words []string) []Stat {
	wordsMap := make(map[string]int)
	for _, v := range words {
		wordsMap[v]++
	}

	counts := make([]Stat, 0, len(wordsMap))
	for word, count := range wordsMap {
		counts = append(counts, Stat{Word: word, Count: count})
	}

	return counts
}

func Sort(a Stat, b Stat) bool {
	if a.Count == b.Count {
		return a.Word < b.Word
	}

	return a.Count > b.Count
}
