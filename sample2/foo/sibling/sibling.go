package sibling

import "sample2/foo/internal"

func AlsoUseDoubler(i int) int {
	return internal.Doubler(i)
}
