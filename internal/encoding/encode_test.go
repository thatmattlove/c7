package encoding_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thatmattlove/c7/internal/encoding"
)

func Test_Encode(t *testing.T) {
	t.Parallel()
	plain := "c7test1"
	encoded, err := encoding.Encode(plain)
	require.NoError(t, err, "failed to encode")
	decoded, err := encoding.Decode(encoded)
	require.NoError(t, err, "failed to decode")
	assert.Equal(t, plain, decoded)
}

func Test_EncodeWithSalt(t *testing.T) {
	for i := 0; i < 8; i++ {
		i := i
		t.Run(fmt.Sprintf("salt-%d", i), func(t *testing.T) {
			t.Parallel()
			plain := fmt.Sprintf("c7test%d", i)
			encoded, err := encoding.EncodeWithSalt(plain, i)
			require.NoError(t, err, "failed to encode")
			decoded, err := encoding.Decode(encoded)
			require.NoError(t, err, "failed to decode")
			assert.Equal(t, plain, decoded)
		})
	}
	t.Run("salt out of range", func(t *testing.T) {
		_, err := encoding.EncodeWithSalt("test", 100)
		require.Error(t, err)
	})
}
