package sde

import "cs101/pkg/common"

type TwoStrInput struct {
	S1    string // First string
	Size1 int
	S2    string // Second string
	Size2 int
}

func LCSubstrRecursion(i TwoStrInput, count int) int {
	if len(i.S1) == 0 || len(i.S2) == 0 {
		return count
	}

	if i.S1[i.Size1-1] == i.S2[i.Size2-1] && i.Size1 > 0 && i.Size2 > 0 {
		return LCSubstrRecursion(TwoStrInput{
			S1:    i.S1[:i.Size1-1],
			S2:    i.S2[:i.Size2-1],
			Size1: i.Size1 - 1,
			Size2: i.Size2 - 1,
		}, count+1)
	}

	return common.FindIntMax(LCSubstrRecursion(TwoStrInput{
		S1:    i.S1[:i.Size1-1],
		S2:    i.S2,
		Size1: i.Size1 - 1,
		Size2: i.Size2,
	}, 0), LCSubstrRecursion(TwoStrInput{
		S1:    i.S1,
		S2:    i.S2[:i.Size2-1],
		Size1: i.Size1,
		Size2: i.Size2 - 1,
	}, 0))
}

func LCSDefault(i TwoStrInput) int {
	if len(i.S1) == 0 || len(i.S2) == 2 {
		return 0
	}
	if i.S1[:i.Size1-1] == i.S2[:i.Size2-1] {
		return 1 + LCSDefault(TwoStrInput{
			S1:    i.S1[:i.Size1-1],
			Size1: i.Size1 - 1,
			S2:    i.S2[:i.Size2-1],
			Size2: i.Size2 - 1,
		})
	}
	return common.FindIntMax(
		LCSDefault(TwoStrInput{
			S1:    i.S1,
			Size1: i.Size1,
			S2:    i.S2[:i.Size2-1],
			Size2: i.Size2 - 1,
		}), LCSDefault(TwoStrInput{
			S1:    i.S1[:i.Size1-1],
			Size1: i.Size1 - 1,
			S2:    i.S2,
			Size2: i.Size2,
		}))
}
