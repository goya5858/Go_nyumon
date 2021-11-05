package mylib

import "testing"

var Debug bool = true

func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("skip Reason") //以下のtestをすべてスキップする
	}
	v := Average([]int{1, 2, 3, 4, 5, 7, 8})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
