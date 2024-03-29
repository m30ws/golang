package main

import (
	"fmt"
	"sync"
)

var (
	wg  sync.WaitGroup
	mux sync.Mutex
	sum int
)

func sumarise(number int) {
	defer wg.Done()
	defer mux.Unlock()
	mux.Lock()
	sum += number
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 1, 2, 3}
	wg.Add(len(numbers))
	for _, number := range numbers {
		num := number
		go sumarise(num)
	}
	wg.Wait()
	fmt.Println(sum) // 21
}