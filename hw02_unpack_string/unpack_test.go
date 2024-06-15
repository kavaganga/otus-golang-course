package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde", err: nil},
		{input: "abccd", expected: "abccd", err: nil},
		{input: "", expected: "", err: nil},
		{input: "aaa0b", expected: "aab", err: nil},
		{input: "a4bc2d5e", expected: "aaaabccddddde", err: nil},
		{input: "abcd", expected: "abcd", err: nil},
		{input: "3abc", expected: "", err: ErrInvalidString},
		{input: "45", expected: "", err: ErrInvalidString},
		{input: "aaa10b", expected: "", err: ErrInvalidString},
		{input: "aaa0b", expected: "aab"},
		{input: "", expected: "", err: nil},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc", err: nil},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			if tc.err != nil {
				require.Error(t, tc.err, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
