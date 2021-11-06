package internal_package_example

//import "sample2/foo/internal"

func FailUseDoubler(i int) int {
	//return internal.Doubler(i)
	return i * 2
}
