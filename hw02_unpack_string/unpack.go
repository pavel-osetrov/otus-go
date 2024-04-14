package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var builder strings.Builder

	var previousValue rune

	for key, value := range str {
		IsNumberValue := IsNumber(value)
		IsNumberPreviousValue := IsNumber(previousValue)
		IsZero := value == '0'

		if key == 0 && IsNumberValue {
			return "", ErrInvalidString
		}

		if IsNumberPreviousValue && IsNumberValue {
			return "", ErrInvalidString
		}

		if IsNumberValue && IsZero {
			str := builder.String()
			updatedStr := str[:len(str)-1]
			builder.Reset()
			builder.WriteString(updatedStr)
		}

		if IsNumberValue && !IsZero {
			count, err := strconv.Atoi(string(value))
			if err != nil {
				return "", err
			}

			repeated := strings.Repeat(string(previousValue), count-1)
			builder.WriteString(repeated)
		}

		if !IsNumberValue {
			builder.WriteString(string(value))
		}
		previousValue = value
	}

	return builder.String(), nil
}

func IsNumber(value rune) bool {
	return '0' <= value && value <= '9'
}
