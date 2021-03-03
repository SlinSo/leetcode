package p0010

import (
	"regexp"
)

func isMatch(s string, p string) bool {
	if !isRegex(p) {
		return s == p
	}

	first_match := s != "" && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (first_match && isMatch(s[1:], p))
	} else {
		return first_match && isMatch(s[1:], p[1:])
	}
}

func isRegex(p string) bool {
	for _, r := range p {
		if r == '.' {
			return true
		}
		if r == '*' {
			return true
		}
	}
	return false
}

func isMatchCheating(s string, p string) bool {
	matched, err := regexp.MatchString("^"+p+"$", s)
	if err != nil {
		return false
	}
	return matched
}
