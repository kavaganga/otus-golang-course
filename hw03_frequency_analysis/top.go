package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(src string) []string {
	var result []string
	data := make(map[string]int)
	nums := 10 // количество элементов для возврата.

	src = strings.ToLower(src)
	r := strings.NewReplacer(
		"\n", " ",
		"\t", " ",
		"!", " ",
		",", " ",
		".", " ",
		"\"", " ",
		"'", " ")
	src = r.Replace(src)

	arr := strings.Split(src, " ")
	for _, w := range arr {
		if w == "" {
			continue
		}
		if w == "-" {
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
		case data[w1] < data[w2]:
			return false
		case data[w1] == data[w2]:
			return w1 < w2
		default:
			return true
		}
	})

	if len(result) < nums {
		return result
	}
	return result[:nums]
}
