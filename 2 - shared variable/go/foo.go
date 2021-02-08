package main

import (
	"fmt"
	"runtime"
)

func number_server(add <-chan int, sub <-chan int, read chan<- int, finished <-chan bool) {
	var number = 0

	// This for-select pattern is one you will become familiar with...
	for {
		select {
		case <- add:
			number++
		case <- sub:
			number--
		case <-finished:
			read <- number
			return 
		}
		
	}
}

func incrementer(add chan<- int, finished chan<- bool) {
	for j := 0; j < 100000; j++ {
		add <- 1
	}
	//signal that the goroutine is finished
	finished <- true
}

func decrementer(sub chan<- int, finished chan<- bool) {
	for j := 0; j < 100000; j++ {
		sub <- 1
	}
	//signal that the goroutine is finished
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Construct the remaining channels
	read := make(chan int)
	add  := make(chan int)
	sub  := make(chan int)
	finished := make(chan bool)
	stop := make(chan bool)


	// Spawn the required goroutines
	go incrementer(add, finished)
	go decrementer(sub, finished)
	go number_server(add, sub, read, stop)
	



	// block on finished from both "worker" goroutines
	<- finished
	<- finished

	stop <- true;

	fmt.Println("The magic number is:", <-read)
}
