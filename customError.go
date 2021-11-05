package main

import "fmt"

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("UserNotFound : %v", e.Username)
}

// 返り値のerrorはError()メソッドを持つインターフェイス
// このケースではError()メソッドを持つ&UserNotFoundが代入されている
func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "Mike"}
}

func main() {
	// myFuncはエラーか否かを返す関数
	// このケースでは絶対エラーを返す(err!=nil)ようになっている
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
