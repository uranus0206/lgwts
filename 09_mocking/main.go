package main

import (
	"lgwts/09_mocking/countdown"
	"os"
)

func main() {
	countdown.Countdown(os.Stdout)
}
