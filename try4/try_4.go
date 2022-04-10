package main

import (
	// myGreeting "greeting/greeting"

	"time"

	greeting "github.com/tenntenn/greeting/v2"
)

func main() {
	// myGreeting.Do()
	println(greeting.Do(time.Now()))
}
