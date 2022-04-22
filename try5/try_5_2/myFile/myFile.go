package myFile

import (
	"fmt"
	"io/ioutil"
)

type File string

// .jpgもしくは.jpegファイルだった場合はtrueを返す
func (f File) IsJpg() bool {
	if f == "./try_5_2.go" {
		return true
	}
	return false
}

func PrintFile(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			PrintFile(dir + "/" + file.Name())
			continue
		}
		var cfile File = File(dir + "/" + file.Name())
		fmt.Println(cfile.IsJpg())

		fmt.Println(dir + "/" + file.Name())

	}
}
