package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)

	files, err := ioutil.ReadDir(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		// if file.IsDir() {
		// 	fmt.Println(file)
		// }
		fmt.Println(file.Name())

	}
	// 指定したディレクトリ以下のJPGファイルをPNGに変換するコマンドを作成する
	// 自作パッケージ
	// ユーザ定義型を作成
	// GoDocを生成
	// Go Modulesを使う

}
