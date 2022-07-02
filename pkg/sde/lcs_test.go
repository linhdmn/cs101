package sde_test

import (
	"testing"

	"cs101/pkg/sde"

	"github.com/stretchr/testify/require"
)

func Test_LCSubstrRecursion(t *testing.T) {
	t.Run("same string should return 0", func(t *testing.T) {
		t.Parallel()
		lcs := sde.LCSubstrRecursion(sde.TwoStrInput{
			S1:    "a",
			Size1: 1,
			S2:    "a",
			Size2: 1,
		}, 0)
		require.Equal(t, 1, lcs)
	})
	t.Run("lcs of ABCDGH and ACDGHR should return 4", func(t *testing.T) {
		t.Parallel()
		lcs := sde.LCSubstrRecursion(sde.TwoStrInput{
			S1:    "ABCDGH",
			Size1: 6,
			S2:    "ACDGHR",
			Size2: 6,
		}, 0)
		require.Equal(t, 4, lcs)
	})
}

func Test_LCSDefault(t *testing.T) {
	t.Run("same string should return 1", func(t *testing.T) {
		t.Parallel()
		lcs := sde.LCSDefault(sde.TwoStrInput{
			S1:    "a",
			Size1: 1,
			S2:    "a",
			Size2: 1,
		})
		require.Equal(t, 1, lcs)
	})
	t.Run("lcs of ABCDGH and ACDGHR should return 4", func(t *testing.T) {
		t.Parallel()
		lcs := sde.LCSDefault(sde.TwoStrInput{
			S1:    "ABCDGH",
			Size1: 6,
			S2:    "ACDGHR",
			Size2: 6,
		})
		require.Equal(t, 4, lcs)
	})
}
