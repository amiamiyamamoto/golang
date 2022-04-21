package myFile

import (
	"fmt"
	"io/ioutil"
)

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
		fmt.Println(dir + "/" + file.Name())

	}
}
