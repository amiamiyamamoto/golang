package main

func main() {
	// var sum float32
	// sum = 5 + 6 + 3
	// avg := sum / 3
	// if avg > 4.5 {
	// 	println("good")
	// }

	// var a, b, c bool
	// a, b, c = true, false, true
	// if a && b || !c {
	// 	println("true\n")
	// } else {
	// 	println("false\n")
	// }

	// type gamePoint struct {
	// 	point int
	// 	user  string
	// }
	// type Score struct {
	// 	UserID string
	// 	GameID int
	// 	Point  int
	// }

	// try3()

	for i := 1; i <= 100; i++ {
		print(i)
		if isOdd(i) {
			println("-奇数")
		} else {
			println("-偶数")
		}
	}

	// n, m := swap(10, 20)
	// println(n, m)
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)

}

func try3() {
	n1 := 19
	n2 := 86
	n3 := 1
	n4 := 12

	sum := n1 + n2 + n3 + n4
	println(sum, "\n")
	////////////////////////////////////////////
	num := []int{19, 86, 1, 12}
	var ans int
	for _, n := range num {
		ans += n
	}
	println(ans)
}

//奇数の場合trueを返す
func isOdd(num int) bool {
	var rsl bool
	if num%2 == 1 {
		rsl = true
	}
	return rsl
}

func swap(n int, m int) (int, int) {
	return m, n
}

func swap2(n *int, m *int) {
	tmp := *n
	*n = *m
	*m = tmp
}
