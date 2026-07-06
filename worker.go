package main

import (
	"fmt"
	"image"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	//fmt.Println("Hello")
	//fmt.Println("Hello")
	urls := []string{"https://google.com", "https://facebook.com",
		"https://x.com", "https://x.com", "https://animepahe.com",
		"https://google.com", "https://google.com", "https://x.com",
		"https://google.com", "http://animepahe.com"}

	// for _, url := range urls {
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println(url)
	// 	}()
	// }
	for _, url := range urls {
		wg.Add(1)
		go func(urlString string) {
			defer wg.Done()
			fmt.Println(urlString)
		}(url)
	}
	wg.Add(1)
	go worker(wg)
	wg.Wait()
}

func worker(wg sync.WaitGroup) {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}

	next := time.After(time.Second)
	for {
		select {
		case <-next:
			wg.Done()
			pos = pos.Add(direction)
			fmt.Println("current location is ", pos)
			next = time.After(time.Second)
			// wait for channel stimuli here.
		}
	}
}
