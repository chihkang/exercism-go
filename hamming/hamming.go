// Package hamming deals with the hamming distance between 2 DNA sequences
package hamming

import (
	"errors"
)

const testVersion = 5

type argError struct {
	arg int
}

// Distance takes in two strings and returns the hamming distance between them
func Distance(a, b string) (int, error) {
	distance := 0

	if a != b {

		alen := len(a)
		blen := len(b)

		if alen != blen {
			return -1, errors.New("Different sizes")
		}

		for i := 0; i < alen; i++ {
			if a[i] != b[i] {
				distance++
			}
		}
	}
	return distance, nil
}
