package mainz

import (
	"fmt"
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
		go func() {
			defer wg.Done()
			fmt.Println(url)
		}()
	}
	//go worker()
	wg.Wait()
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
