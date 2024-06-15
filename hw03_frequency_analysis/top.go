package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(src string) []string {
	var result []string
	data := make(map[string]int)
	arr := strings.Split(src, " ")
	for _, w := range arr {
		if w == "" {
			continue
		}
		data[w]++
	}
	for k := range data {
		result = append(result, k)
	}

	sort.Slice(result, func(i, j int) bool {
		w1 := result[i]
		w2 := result[j]
		switch {
		case data[w1] > data[w2]:
			return true
		case data[w1] == data[w2]:
			return w1 > w2
		default:
			return false
		}
	})
	if len(result) < 10 {
		return result
	}
	return result[:10]
}
