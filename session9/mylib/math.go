package mylib

func Average(s []int) int {
	total := 0
	for _, num := range s {
		total += num
	}
	return total / len(s)
}
