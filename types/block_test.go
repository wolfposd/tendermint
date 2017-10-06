package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlockIDCopy(t *testing.T) {
	sHash := "0xdeadbeef"
	bi := &BlockID{
		Hash: []byte("Foo-Tendermint"),

		PartsHeader: PartSetHeader{
			Hash: []byte(sHash),
		},
	}

	bic := bi.Copy()
	require.Equal(t, bi, bic, "BlockID.Copy should return the same")

	bic.PartsHeader.Hash[2] = 0x61
	require.NotEqual(t, bi, bic, "a mutation of the copy should not mutate the original")

	bic.Hash[2] = 0x41
	require.NotEqual(t, bi, bic, "a mutation of the copy should not mutate the original")
}
