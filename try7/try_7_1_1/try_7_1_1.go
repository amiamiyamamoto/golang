package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

func main() {
	// intのiを定義
	var i int = 10
	// ToStringerメソッドでiをStringer型に変換
	si, err := ToStringer(i)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(si.String())

}

// 任意の値をStringer型に変換する関数
func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, MyError("CastError")
}

// error interefaceを実装したユーザ定義型
type MyError string

func (e MyError) Error() string {
	return string(e)
}
