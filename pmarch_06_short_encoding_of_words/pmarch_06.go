package pmarch_06

import (
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {
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
