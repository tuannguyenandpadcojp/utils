package base62_test

import (
	"fmt"
	"testing"

	"github.com/tuannguyenandpadcojp/utils/base62"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		decimal int64
		base62  string
	}{
		{0, "0"},
		{1, "1"},
		{-1, "-1"},
		{61, "z"},
		{-61, "-z"},
		{62, "10"},
		{-62, "-10"},
		{62 + 1, "11"},
		{62 + 61, "1z"},
		{62 * 62, "100"},
		{62*62 - 1, "zz"},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%d -> %s", tt.decimal, tt.base62)
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := base62.Encode(tt.decimal); got != tt.base62 {
				t.Errorf("Encode() = %v, want %v", got, tt.base62)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		base62  string
		decimal int64
		invalid bool
	}{
		// invalid base62 string
		{"", 0, true},
		{"@", 0, true},
		{"-", 0, true},
		{"1-", 0, true},

		// valid base62 string
		{"0", 0, false},
		{"1", 1, false},
		{"-1", -1, false},
		{"z", 61, false},
		{"-z", -61, false},
		{"10", 62, false},
		{"-10", -62, false},
		{"11", 62 + 1, false},
		{"1z", 62 + 61, false},
		{"100", 62 * 62, false},
		{"zz", 62*62 - 1, false},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s -> %d", tt.base62, tt.decimal)
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := base62.Decode(tt.base62)
			if tt.invalid && (err == nil) {
				t.Errorf("Decode() should return an error but got nil")
				return
			}
			if got != tt.decimal {
				t.Errorf("Decode() = %v, want %v", got, tt.decimal)
			}
		})
	}
}

func BenchmarkEncodePositiveNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		base62.Encode(int64(-i))
	}
}

func BenchmarkEncodeNegativeNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		base62.Encode(int64(-i))
	}
}

func BenchmarkEncodeDecodePositiveNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = base62.Decode(base62.Encode(int64(i)))
	}
}

func BenchmarkEncodeDecodeNegativeNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = base62.Decode(base62.Encode(int64(-i)))
	}
}
