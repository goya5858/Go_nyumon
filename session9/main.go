package main

import (
	"fmt"
	//以下エラーでてるけど読み込める
	//"mymod/mylib"
	a "github.com/markcheno/go-quote"
	b "github.com/markcheno/go-talib"
	//"mymod/mylib/under"
	// mymodが自分で明示的につけたモジュール名
	// >> go mod init mymod
	// mypacが自分が明示的につけたパッケージ名
)

func main() {
	//s := []int{1, 2, 3, 4, 5}
	//fmt.Println(mylib.Average(s))
	//mylib.Say()
	//person := mylib.person{Name: "Mike", Age: 20}
	//fmt.Println(person)

	spy, _ := a.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", a.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := b.Rsi(spy.Close, 2)
	fmt.Println(rsi2)

}
