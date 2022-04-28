package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// readLineByScanner()
	// readWordByScanner()
	// readCustomScanner()
	readCustomScannerCommaSeparated()
}

// デフォrとのsplitFuncを使用して行単位で読み込み
func readLineByScanner() {
	reader := strings.NewReader("1234\n5678")
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 標準で用意されているScanWords関数を利用して、単語単位で読み込み
func readWordByScanner() {
	reader := strings.NewReader("abc efg\nhijk lmn")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 独自のsplitFuncを定義して、数値を読み込み
// 数値として変換できない場合はエラー
func readCustomScanner() {
	scanner := bufio.NewScanner(strings.NewReader("1234 5678 1234567891234567890"))

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}

	scanner.Split(split)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

// 独自のsplitFunc関数を定義して、カンマで区切られている単語ごとに読み込み
func readCustomScannerCommaSeparated() {
	scanner := bufio.NewScanner(strings.NewReader("1,2,3,4"))
	onComma := func(data []byte, atEOf bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		return 0, data, bufio.ErrFinalToken
	}

	scanner.Split(onComma)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
