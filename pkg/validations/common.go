package validations

import "regexp"

// Custom function to match regex
func RegexMatch(value string, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
}
