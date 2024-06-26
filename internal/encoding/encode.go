package encoding

import (
	"fmt"
	"math/rand/v2"
)

func EncodeWithSalt(plain string, salt int) (string, error) {
	s := salt
	if s > 7 {
		return "", fmt.Errorf("salt must be in range 0..7, got %d", salt)
	}
	pb := []byte(plain)
	encoded := fmt.Sprintf("%.2o", s)
	for i := 0; i < len(pb); i++ {
		r := int(pb[i])
		x := int(XLAT[s])
		b := r ^ x
		encoded += fmt.Sprintf("%.2X", b)
		s++
	}
	return encoded, nil
}

func Encode(plain string) (string, error) {
	s := rand.IntN(7)
	return EncodeWithSalt(plain, s)
}
