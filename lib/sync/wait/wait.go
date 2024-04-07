package wait

import (
	"sync"
	"time"
)

type Wait struct {
	wg sync.WaitGroup
}

func (w *Wait) Done() {
	w.wg.Done()
}

func (w *Wait) Add(delta int) {
	w.wg.Add(delta)
}

func (w *Wait) Wait() {
	w.wg.Wait()
}

func (w *Wait) WaitIfTimeOut(t time.Duration) bool {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		w.Wait()
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		return false
	case <-time.After(t):
		return true
	}
}
