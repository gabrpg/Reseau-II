package main

import "sync"

type WaitGroup interface {
	Add(delta int)
	Done()
	Wait()
}

type SemaphoreWaitGroup struct {
	sem chan bool
	wg  sync.WaitGroup
}

func (s *SemaphoreWaitGroup) Add(delta int) {
	s.wg.Add(delta)
	s.sem <- true
}

func (s *SemaphoreWaitGroup) Done() {
	<-s.sem
	s.wg.Done()
}

func (s *SemaphoreWaitGroup) Wait() {
	s.wg.Wait()
}
