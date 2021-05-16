package main

import (
	"fmt"
	"sync"
)

// sync + chan 控制并发

type Store struct {
	items map[int]struct{}
	lock  sync.RWMutex
}

func (s *Store) Add(key int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items[key] = struct{}{}
}

func main() {
	count := 1000

	var wg sync.WaitGroup
	s := &Store{
		items: make(map[int]struct{}),
	}

	wg.Add(count)

	ch := make(chan struct{}, 200)

	for i := 0; i < count; i++ {
		ch <- struct{}{}

		go func(i int, s *Store) {
			defer wg.Done()
			s.Add(i)
			<-ch
		}(i, s)
	}

	wg.Wait()

	fmt.Println("legth", len(s.items))
}
