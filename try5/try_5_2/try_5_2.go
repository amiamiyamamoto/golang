package main

import (
	"flag"
	myfile "myfile/myFile"
)

func main() {
	flag.Parse()
	// args := flag.Args()

	myfile.ChangeFilesJpgToPng(flag.Args()[0])

}
