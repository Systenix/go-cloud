package generators

import "strings"

// lowerFirst converts the first character of a string to lowercase.
func lowerFirst(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToLower(str[:1]) + str[1:]
}

// toCamelCase converts the str to camelCase and upper first letter.
func toCamelCase(str string) string {
	words := strings.Split(str, "_")
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}
	return strings.Join(words, "")
}
