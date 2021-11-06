package foo

import "sample2/foo/internal"

func UseDoubler(i int) int {
	return internal.Doubler(i)
}
