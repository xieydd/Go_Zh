package main

import (
	"fmt"
)

func sum(values []int, channels chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	channels <- sum
}

func main() {
	var values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	channels := make(chan int, 2)
	go sum(values[:len(values)/2], channels)
	go sum(values[len(values)/2:], channels)
	sum1, sum2 := <-channels, <-channels
	fmt.Println("sum1: ", sum1, "  sum2: ", sum2)
}
