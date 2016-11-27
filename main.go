package main

import (
	"fmt"

	"github.com/shogo-ma/golang-algorithm-samples/sort"
)

func main() {

	// selection sort
	fmt.Println([]int{4, 5, 1, 3, 8, 2})
	fmt.Println(sort.Selection([]int{4, 5, 1, 3, 8, 2}))

	// insertion sort
	fmt.Println([]int{4, 5, 1, 3, 8, 2})
	fmt.Println(sort.Insertion([]int{4, 5, 1, 3, 8, 2}))

	// merge sort
	fmt.Println([]int{4, 5, 1, 3, 8, 2})
	fmt.Println(sort.Merge([]int{4, 5, 1, 3, 8, 2}, 0, 5))
}
