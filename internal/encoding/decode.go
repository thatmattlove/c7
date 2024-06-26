package encoding

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func Decode(encoded string) (string, error) {
	pattern := regexp.MustCompile("(^[0-9A-Fa-f]{2})([0-9A-Fa-f]+)")
	if !pattern.MatchString(encoded) {
		return "", fmt.Errorf("invalid encoded value")
	}
	result := pattern.FindAllStringSubmatch(encoded, -1)
	rm := map[int]string{}
	for i, n := range result[0] {
		rm[i] = n
	}
	// Create an int from the first two characters of the string.
	r1, err := strconv.ParseInt(rm[1], 10, 0)
	if err != nil {
		return "", errors.Join(fmt.Errorf("failed to parse encoded value"), err)
	}
	// Get the remainder of the string after the first two characters.
	r2 := rm[2]

	decoded := ""

	// Iterate through the 2nd part of the string.
	for pos := 0; pos < len(r2); pos++ {
		// Skip odd indices.
		if pos%2 == 0 {
			// Stringify this character.
			m1 := string(r2[pos])
			// Stringify the next character.
			m2 := string(r2[pos+1])
			// Convert the resulting combined strings to base16 integer.
			m, err := strconv.ParseInt(m1+m2, 16, 0)
			if err != nil {
				return "", errors.Join(fmt.Errorf("failed to parse encoded value at position %d", pos), err)
			}
			var newchar string
			if r1 <= 50 {
				idx := byte(m) ^ XLAT[r1]
				newchar = fmt.Sprintf("%c", idx)
				r1 += 1
			}
			if r1 == 51 {
				r1 = 0
			}
			decoded += newchar
		}
	}
	return decoded, nil
}
