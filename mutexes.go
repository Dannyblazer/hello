package main

import (
	"fmt"
	"sync"
	"time"
)

// Visited tracks whether web pages have been visited
// Its method may be used concurrently for multiple goroutines
//type upstream chan

type Visited struct {
	// mu guards the visited maps
	mu      sync.Mutex
	visited map[string]int
}

func main() {
	upstream := make(chan string)

	//downstream := make(chan string)
	go source(upstream)
	go visit(upstream)
	time.Sleep(3 * time.Second)

}

func source(c chan string) {
	urls := []string{"https://google.com", "https://facebook.com",
		"https://x.com", "https://x.com", "https://animepahe.com",
		"https://google.com", "https://google.com", "https://x.com",
		"https://google.com", "http://animepahe.com"}
	for _, url := range urls {
		c <- url
	}
	close(c)
}

func visit(c chan string) {
	visit := Visited{visited: map[string]int{}}
	for item := range c {
		visit.VisitLink(item)
	}
	fmt.Println(visit.visited)
}

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}
