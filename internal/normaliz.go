package internal

import (
	"fmt"
	"strings"
	"vardoc/internal/assert"
)

var (
	ErrInvalidForamt error = fmt.Errorf("Invalid code format")
)

// Normalizes a golang code snipped into a standarized form.
// If the code is not valid golang this function returns a ErrInvalidForamt
func NormalizGolangCodeSnipped(codeSnipped string) (string, error) {
	var (
		normalizedCodeSnipped strings.Builder = strings.Builder{}
		startIdx              int             = 0
		codeInCurrentLine     bool            = false
		previousSpace         bool            = false
		previousNewLine       bool            = false
	)

	// skip first till first code
	for idx, r := range codeSnipped {
		switch r {
		case ' ', '\n', '\t':
			startIdx = idx
			continue
		}

		break
	}
	codeSnipped = codeSnipped[startIdx:]

	for i := 0; i < len(codeSnipped); i++ {
		r := codeSnipped[i]

		if r == '"' {
			err := normalizedCodeSnipped.WriteByte(r)
			assert.NoError(err)
			i += 1

			ignoreNext := false
			for ignoreNext || codeSnipped[i] != '"' {
				r = codeSnipped[i]

				hasNextRune := len(codeSnipped) > i+1
				if hasNextRune == false {
					return "", ErrInvalidForamt
				}

				ignoreNext = (r == byte('\\'))

				err := normalizedCodeSnipped.WriteByte(r)
				assert.NoError(err)

				i += 1
			}

			r = codeSnipped[i]
			err = normalizedCodeSnipped.WriteByte(r)
			assert.NoError(err)

			i += 1

			continue
		}

		switch r {
		case ' ':
			if codeInCurrentLine == false || previousSpace {
				continue
			}

			hasNextRune := len(codeSnipped) > i+1
			if hasNextRune == false {
				continue
			}

			nextRune := codeSnipped[i+1]
			if nextRune == '\n' || nextRune == ' ' || nextRune == '\t' {
				continue
			}

			previousSpace = true

		case '\n':
			if previousNewLine {
				continue
			}
			previousNewLine = true
			codeInCurrentLine = false

		case '\t':
			continue

		default:
			previousSpace = false
			codeInCurrentLine = true

		}

		err := normalizedCodeSnipped.WriteByte(r)
		assert.NoError(err)
	}

	return normalizedCodeSnipped.String(), nil
}
