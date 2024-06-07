package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{300, 200, 100} {
			c <- num
		}
		close(c)
	}()

	for num := range joinChannels(a, b, c) {
		fmt.Println(num)
	}

}

func joinChannels(channels ...<-chan int) chan int {
	result := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	go func() {
		defer close(result)

		for _, ch := range channels {
			go func() {
				defer wg.Done()
				for num := range ch {
					result <- num
				}
			}()
		}

		wg.Wait()
	}()

	return result
}
