package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func Unpack(src string) (string, error) {
	dst := ""
	prevNumber := false
	prevSlash := false
	for i, ch := range src {
		switch {
		case IsDigit(ch):
			{
				if prevNumber || i == 0 {
					return "", ErrInvalidString
				}

				if !prevSlash {
					prevNumber = true
				}
				n, _ := strconv.Atoi(string(ch))
				if n < 0 {
					return "", ErrInvalidString
				}
				switch {
				case n == 0:
					{
						dst = dst[:len(dst)-1]
					}
				case prevSlash:
					{
						prevNumber = false
						prevSlash = false
						dst += string(ch)
					}
				default:
					{
						dst += strings.Repeat(string(src[i-1]), n-1)
					}
				}
			}
		case prevSlash && ch == '\\':
			{
				dst += string(ch)
				prevSlash = false
			}
		case ch == '\\':
			{
				prevSlash = true
			}
		default:
			dst += string(ch)
			prevNumber = false
		}
	}
	return dst, nil
}
