package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader("123\n456"))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
