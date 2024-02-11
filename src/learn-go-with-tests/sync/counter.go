package main

import "sync"

type Counter struct {
	mu sync.Mutex
	value int
}

func (c *Counter) Inc() {
	// Incを呼び出すときにロックをかける
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}