package global

import "sync"

type Global struct {
	mu sync.Mutex
	count int
}

func (g *Global) GetCount() int {
	return g.count
}

func (g *Global) SetIncrement() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.count++
}