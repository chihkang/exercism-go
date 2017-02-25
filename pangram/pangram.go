package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(input string) bool {

	for _, ch := range "abcdefghijklmnopqrstuvwxyz" {
		if strings.Contains(input, string(ch)) || strings.Contains(input, strings.ToUpper(string(ch))) {
			continue
		}
		return false
	}
	return true
}
// BenchmarkPangram-4   	  500000	      3490 ns/op
