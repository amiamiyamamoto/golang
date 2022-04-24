package main

import (
	"flag"
	"fmt"
	myfile "myfile/myFile"
)

func main() {
	flag.Parse()
	// args := flag.Args()

	fmt.Println(myfile.ChangeFilesJpgToPng(flag.Args()[0]))

}
