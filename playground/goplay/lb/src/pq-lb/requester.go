package main

import (
	"log"
	"math/rand"
	"time"
)

var counter int

type Request struct {
	fn func() int // operation to perform
	c  chan int   // channel on which to return result
}

func requester(work chan Request) {
	c := make(chan int, 100)
	go func() {
		for {
			time.Sleep(time.Microsecond * time.Duration(rand.Int63n(10))) // simulate uneven throughput
			work <- Request{do_some_work, c}                              // send a work request
			result := <-c                                                 // wait for answer
			do_something_else(result)
		}
	}()
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func do_some_work() int {
	//defer trace("do_some_work")()
	counter++
	time.Sleep(time.Microsecond * time.Duration(rand.Int63n(10))) // simulate some actual work.
	return counter
}

func do_something_else(r int) {
	// do very important work.
}
