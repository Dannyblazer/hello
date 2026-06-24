package main

import (
	"fmt"
	"sync"
)

type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func main() {
	upstream := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go source(upstream)
	go func() {
		defer wg.Done()
		visit(upstream)
	}()

	wg.Wait()
}

func source(c chan string) {
	urls := []string{
		"https://google.com", "https://facebook.com",
		"https://x.com", "https://x.com", "https://animepahe.com",
		"https://google.com", "https://google.com", "https://x.com",
		"https://google.com", "http://animepahe.com",
	}
	for _, url := range urls {
		c <- url
	}
	close(c)
}

func visit(c chan string) {
	v := Visited{visited: map[string]int{}}
	for item := range c {
		v.VisitLink(item)
	}
	fmt.Println(v.visited)
}

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}
