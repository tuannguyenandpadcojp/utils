// Package base62 provides functions to encode and decode base62 strings.
// Reference: https://en.wikipedia.org/wiki/Base62
package base62

import (
	"bytes"
	"fmt"
)

var charset = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// Encode encodes a decimal number to base62 string.
func Encode(n int64) string {
	if n == 0 {
		return string(charset[0])
	}

	if n < 0 {
		return "-" + Encode(-n)
	}

	var result []byte
	for n > 0 {
		result = prependByte(result, charset[n%62])
		n /= 62
	}
	return string(result)
}

func prependByte(x []byte, y byte) []byte {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}

// Decode decodes a base62 string to decimal number.
func Decode(s string) (int64, error) {
	// special invalid cases
	if s == "" || s == "-" {
		return 0, fmt.Errorf("the input string '%s' is invalid", s)
	}

	var n int64
	for i, c := range s {
		// first character can be '-' for negative number
		if i == 0 && c == '-' {
			continue
		}
		idx := bytes.IndexRune(charset, c)
		if idx == -1 {
			return 0, fmt.Errorf("the input string '%s' contains invalid character: '%s' at position %d", s, string(c), i)
		}
		n = n*62 + int64(idx)
	}

	// negative number
	if s[0] == '-' {
		n = -n
	}
	return n, nil
}
