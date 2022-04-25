package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Stringer interface {
	String() string
}

// ユーザ定義型のStringerInt型にStringメソッドを実装
type StringerInt int

func (s StringerInt) String() string {
	return strconv.Itoa(int(s))
}

// floatでStringer interfaceを実装する
type StringerFloat64 float64

func (s StringerFloat64) String() string {
	return strconv.FormatFloat(float64(s), 'f', -1, 64)
}

func main() {
	var i StringerInt = -10
	var f StringerFloat64 = 12.34

	// Stringerのスライスを作成し、ループする
	var list []Stringer = []Stringer{i, f}
	for _, i := range list {
		fmt.Println(getType(i))
	}

}

// Stringer interface型を受け取り、具象型名と値を返す
func getType(s Stringer) string {
	type_and_value := "type: " + reflect.TypeOf(s).Name() + ", value: " + s.String()
	return type_and_value
}
