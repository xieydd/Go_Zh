package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	left, right := 0, len(array)-1
	pivot := rand.Int() % len(array)
	array[pivot], array[right] = array[right], array[pivot]

	for i := range array {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}
	array[left], array[right] = array[right], array[left]
	QuickSort(array[:left])
	QuickSort(array[left+1:])
	return array
}

func Modify(array [5]int) [5]int {
	array[0] = 10
	return array
}

func InverArray(array []int) []int {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func main() {
	slices := []int{100, 2, 3, 4, 5}
	arrays := [5]int{100, 2, 3, 4, 5}
	fmt.Println("Origin array is ", slices)
	result := QuickSort(slices)
	result1 := Modify(arrays)
	result2 := InverArray(slices)
	fmt.Printf("Origin arrays is %v\n", arrays)
	fmt.Printf("After Modify, arrays is %v\n", result1)
	fmt.Printf("After QuickSort Origin slices address is %v\n", &slices[0])
	fmt.Printf("After QuickSort, Result slices address is %v\n", &result[0])
	fmt.Println("After QuickSort, Result slices is ", result)
	fmt.Println("After InverArray, Result slices is ", result2)
}
