package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()        //標準入力を受け取る
	return sc.Text() //最新の受け取った入力をStringで返す
}

// １文字(1数字)毎にintに変換し、一次元配列に挿入していく
func make_intline() []int {
	var intline []int
	inputs := strings.Split(nextLine(), " ") //" "で区切られた一行の文字列を []string にSplitする
	for _, inp := range inputs {
		inp_num, _ := strconv.Atoi(inp)
		intline = append(intline, inp_num)
	}
	return intline
}

func makeMatrix(num int) [][]int {
	var matrix [][]int //空のint二次元スライスを作成
	for i := 0; i < num; i++ {
		intline := make_intline()
		matrix = append(matrix, intline)
	}
	return matrix
}

func main() {
	num, _ := strconv.Atoi(nextLine()) // Stringで入力されたNを int, err に変換する
	matrix := makeMatrix(num)
	fmt.Printf("%v\n", matrix)
}
