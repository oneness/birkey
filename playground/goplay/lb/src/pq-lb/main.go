package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	procs := runtime.NumCPU()
	runtime.GOMAXPROCS(procs)
	fmt.Printf("\nRunning on %d processors.\n", procs)
	nworkers := procs - 2

	rand.Seed(8)

	work := make(chan Request, 100)

	balancer := new_balancer(nworkers, work)
	balancer.start()
	balancer.balance(work)

	for i := 1; i <= 100; i++ {
		requester(work)
	}

	go func() {
		for _ = range time.Tick(250 * time.Millisecond) {
			balancer.print() // periodically print out the number of pending tasks assigned to each worker.
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Printf("\n %d jobs complete.\n", counter)
}
