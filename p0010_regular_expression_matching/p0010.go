package p0010

import (
	"regexp"
)

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	if p == "" {
		return s == ""
	}

	first_match := s != "" && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (first_match && isMatch(s[1:], p))
	} else {
		return first_match && isMatch(s[1:], p[1:])
	}
}

func isMatchCheating(s string, p string) bool {
	matched, err := regexp.MatchString("^"+p+"$", s)
	if err != nil {
		return false
	}
	return matched
}
