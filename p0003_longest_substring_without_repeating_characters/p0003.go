package p0003_longest_substring

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	for i := range s {
		m := make(map[byte]bool)

		m[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			if m[s[j]] {
				break
			}
			m[s[j]] = true
		}

		if len(m) > maxLen {
			maxLen = len(m)
			if maxLen >= len(s)-i {
				break
			}
		}
	}

	return maxLen
}

int n = s.length(), ans = 0;
Map<Character, Integer> map = new HashMap<>(); // current index of character
// try to extend the range [i, j]
for (int j = 0, i = 0; j < n; j++) {
	if (map.containsKey(s.charAt(j))) {
		i = Math.max(map.get(s.charAt(j)), i);
	}
	ans = Math.max(ans, j - i + 1);
	map.put(s.charAt(j), j + 1);
}
return ans;