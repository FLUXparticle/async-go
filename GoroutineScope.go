package async

import "sync"

type GoroutineScope struct {
	wg *sync.WaitGroup
}

func RunBlocking(block func(g *GoroutineScope)) {
	var wg sync.WaitGroup
	block(&GoroutineScope{wg: &wg})
	wg.Wait()
}

func (g *GoroutineScope) Async(block func()) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		block()
	}()
}
