package encoding_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thatmattlove/c7/internal/encoding"
)

func Test_Decode(t *testing.T) {
	t.Run("value1", func(t *testing.T) {
		result, err := encoding.Decode("121A520317181855")
		require.NoError(t, err)
		assert.Equal(t, "c7test1", result)
	})
	t.Run("error", func(t *testing.T) {
		_, err := encoding.Decode(";ljksjsf89f8fkskf")
		assert.Error(t, err)
	})
}
