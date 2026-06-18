package main

import (
	"fmt"
	"sync"
)

// Visited tracks whether web pages have been visited
// Its method may be used concurrently for multiple goroutines

type Visted struct {
	// mu guards the visited maps
	mu      sync.Mutex
	visited map[string]int
}

func main() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Hello")
}

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}
