package p0820

import (
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	if len(words) == 0 {
		return 0
	}

	m := make(map[string]struct{})
	for _, w := range words {
		m[w] = struct{}{}
	}

	for _, w := range words {
		for j := 1; j < len(w); j++ {
			delete(m, w[j:])
		}
	}

	ans := 0
	for k := range m {
		ans += len(k) + 1
	}

	return ans
}

func minimumLengthEncodingSimple(words []string) int {
	if len(words) == 0 {
		return 0
	}

	sort.Slice(words, func(i, j int) bool { return len(words[i]) > len(words[j]) })

	s := ""
	for _, w := range words {
		i := strings.Index(s, w+"#")
		if i == -1 {
			s += w + "#"
		}
	}

	return len(s)
}
