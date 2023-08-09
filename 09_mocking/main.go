package main

import (
	"lgwts/09_mocking/countdown"
	"os"
	"time"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{Duration: 1000 * time.Millisecond}
	countdown.Countdown(os.Stdout, sleeper)
}
