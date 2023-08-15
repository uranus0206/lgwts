package counter

import "sync"

func NewCounter() *Counter {
	return &Counter{}
}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() (count int) {
	return c.value
}
