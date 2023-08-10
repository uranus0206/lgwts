package countdown

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWorld     = "Go!"
	countdownStart = 3
	sleep          = "sleep"
	write          = "write"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.Duration)
}

type OperationCountdownSpy struct {
	Calls []string
}

func (c *OperationCountdownSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}
func (c *OperationCountdownSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	_ = p // fix revive warning
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	// for i := countdownStart; i > 0; i-- {
	// 	fmt.Fprintln(out, i)
	// }
	sleeper.Sleep()
	fmt.Fprint(out, finalWorld)
}
