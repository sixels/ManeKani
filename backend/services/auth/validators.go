package auth

import "regexp"

const (
	usernameMinLen = 4
	usernameMaxLen = 20
)

var usernameRules = []*regexp.Regexp{
	// limit valid characters
	regexp.MustCompile(`^[a-zA-Z0-9._-]+$`),
	// starts with a letter
	regexp.MustCompile(`^[a-zA-Z]`),
	// not end with a symbol
	regexp.MustCompile(`[a-zA-Z0-9]$`),
	// no consecutive symbols
	regexp.MustCompile(`^([a-zA-Z0-9]|[._-][a-zA-Z0-9])*$`),
}

func ValidateUsername(username string) bool {
	usernameLen := len(username)
	return usernameLen >= usernameMinLen && usernameLen <= usernameMaxLen && matchRules(usernameRules, username)
}

func matchRules(rules []*regexp.Regexp, s string) bool {
	for _, rule := range rules {
		if !rule.MatchString(s) {
			return false
		}
	}
	return true
}
