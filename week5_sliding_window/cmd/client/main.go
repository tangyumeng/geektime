package main

import (
	"math/rand"
	"sync"
	"time"

	"example.com/hystrix-demo/testcase"
)

func main() {
	rand.Seed(time.Now().Unix())

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			testcase.Client()
			wg.Done()
		}()
	}

	wg.Wait()
}
