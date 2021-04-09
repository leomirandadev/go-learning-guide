package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runAndWaitRoutines()
	go fmt.Println("Executed")
}

func runAndWaitRoutines() {
	startIn := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ { // with runtime.NumCPU() you can know how many cores your computer have
		wg.Add(1)
		go sayHello(i, &wg)
	}

	wg.Wait()

	fmt.Println("Executed in", time.Now().Sub(startIn))
}

func sayHello(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second * 2)
	fmt.Println("Go routine", i)
}
