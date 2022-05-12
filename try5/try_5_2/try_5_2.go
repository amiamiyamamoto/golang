package main

import (
	"flag"
	"fmt"
	myfile "myfile/myFile"
	"os"
)

func main() {
	flag.Parse()
	// args := flag.Args()

	fmt.Println(myfile.ChangeFilesJpgToPng(flag.Args()[0]))
	if err := myfile.ChangeFilesJpgToPng(flag.Args()[0]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}
