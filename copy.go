package main

import (
	"fmt"
	"sync"
)

func main() {
	orders := make(chan string)
	var wg sync.WaitGroup

	// go func() {
	// 	for i := 0; i <= 5; i++ {
	// 		order := fmt.Sprintf("Order %d", i)
	// 		orders <- order
	// 		fmt.Println("Order placed")
	// 	}
	// 	close(orders)
	// }()

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			order := fmt.Sprintf("Order %d", i)
			orders <- order
			fmt.Printf("%s placed\n", order)
		}(i)
	}

	go func() {
		wg.Wait()
		close(orders)

	}()

	for order := range orders {
		fmt.Println("Delivered ", order)
	}
}
