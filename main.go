package main

import (
	"fmt"
	"github.com/Seven11Eleven/go-pod-kapot/mutex_impl"
	"sync"
)

const goCount = 5555555

func main() {
	myMut := mutex_impl.NewMutexWithChan()
	cnt := 0
	wg := sync.WaitGroup{}

	wg.Add(goCount)

	for i := 0; i < goCount; i++ {
		go func() {
			defer wg.Done()
			myMut.Lock()
			cnt++
			myMut.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Final count:", cnt)
}
