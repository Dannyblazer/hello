package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	stream := make(chan string)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(chefID int) {
			defer wg.Done()
			for j := 1; j <= 4; j++ {
				meal := fmt.Sprintf("Chef %d - Meal %d", chefID, j)
				stream <- meal
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(stream)
	}()

	for meal := range stream {
		fmt.Println("Served", meal)
	}
}
