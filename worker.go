package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("Hello")
	fmt.Println("hello there")
	go worker()
}

func worker() {
	n := 0
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			n++
			fmt.Println(n)
			next = time.After(time.Second)
			// wait for channel stimuli here.
		}
	}
}
