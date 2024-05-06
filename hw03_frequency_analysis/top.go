package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func Top10(text string) []string {
	wordFreq := countWords(text)

	return topNWords(wordFreq, 10)
}

// подсчитывает частоту слов в тексте.
func countWords(text string) map[string]int {
	wordFreq := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		wordFreq[word]++
	}

	return wordFreq
}

// возвращает топ N самых частотных слов.
func topNWords(wordFreq map[string]int, n int) []string {
	counts := make([]wordCount, 0, len(wordFreq))

	for word, count := range wordFreq {
		counts = append(counts, wordCount{word, count})
	}

	sortWordCounts(counts)

	if len(counts) > n {
		counts = counts[:n]
	}

	result := make([]string, len(counts))

	for i, wc := range counts {
		result[i] = wc.word
	}

	return result
}

// сортирует слайс wordCount по убыванию частоты слов, а при равенстве - по алфавиту.
func sortWordCounts(counts []wordCount) {
	sort.Slice(counts, func(i, j int) bool {
		if counts[i].count == counts[j].count {
			return counts[i].word < counts[j].word
		}
		return counts[i].count > counts[j].count
	})
}
