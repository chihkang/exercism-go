package acronym

import (
	"strings"
)

const testVersion = 2

func Abbreviate(input string) (result string) {
	if strings.Contains(input, ":") {
		result = input[:strings.Index(input, ":")]
		return
	}

	f := func(c rune) bool {
		switch c {
		case '-', ' ':
			return true
		}
		return false
	}

	words := strings.FieldsFunc(input, f)
	for _, word := range words {
		for i := range word {
			if i == 0 {
				result += string(word[i])
			} else if word[i] >= 'A' && word[i] <= 'Z' {
				result += string(word[i])
			}
		}
	}
	result = strings.ToUpper(result)

	return result
}
