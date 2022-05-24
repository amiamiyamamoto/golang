package main

import (
	"fmt"
	"os"
)

func main() {

	url := "https://~~~~"
	fn := "sample.jpg"
	if err := download(url, fn); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func dowmload(url string, fn string) error {

	return nil
}
