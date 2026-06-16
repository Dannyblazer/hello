package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Sending channel message")
	c := make(chan int)

	upstream := make(chan string)
	downstream := make(chan string)

	fmt.Println("sleeping....")

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	// for i := 0; i < 5; i++ {
	// 	gopherID := <-c
	// 	fmt.Println("gopher ", gopherID, "has finished sleeping")
	// }
	time.Sleep(4 * time.Second)

	// Channel surfing
	timeout := time.After(2 * time.Second)

	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gophers ", gopherID, "has finished sleeping")
		case <-timeout:
			fmt.Println("Ran out of patience")
			return
		}
	}

	// Pipeline code
	fmt.Println("Starting Gopher Pipeline")

	go sourceGopher(downstream)
	go filterGopher(downstream, upstream)

	printGopher(upstream)

	// Experiment: remove-identical.go

}

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("...", id, "snore ...")
	c <- id // send first
	if id == 3 {
		time.Sleep(10 * time.Second) // block AFTER sending
		fmt.Println("gopher 3 finally done")
	}
}

func sourceGopher(downstream chan string) {
	source := []string{"first", "second", "third", "fourth", "fifth", "fifth"}
	for _, data := range source {
		downstream <- data
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	prev := ""
	for item := range upstream {
		if strings.Contains(item, "bad") {
			continue
		}
		if item != prev {
			downstream <- item
			prev = item
		}
	}
	close(downstream)

}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
