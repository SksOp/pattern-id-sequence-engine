package utils

import "strings"

type Pattern struct {
	pattern string //"YYYYMMDDIIIICC"
}

func NewPattern(pattern string) *Pattern {
	return &Pattern{
		pattern: pattern,
	}
}

func ParsePattern(input string) []string {

	var result []string
	var current strings.Builder
	var prevChar rune

	for i, char := range input {
		if i == 0 {
			current.WriteRune(char)
			prevChar = char
			continue
		}
		if char == prevChar {
			current.WriteRune(char)
		} else {
			result = append(result, current.String())
			current.Reset()
			current.WriteRune(char)
		}
		prevChar = char
	}
	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}
