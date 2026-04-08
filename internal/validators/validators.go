package validators

import "strings"

func NonEmpty(s string) bool {
	return strings.TrimSpace(s) != ""
}

func EmailLooksOK(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) < 5 || len(s) > 254 {
		return false
	}
	at := strings.Index(s, "@")
	if at < 1 || at == len(s)-1 {
		return false
	}
	return strings.Contains(s[at+1:], ".")
}

func RatingOK(r int) bool {
	return r >= 1 && r <= 5
}
