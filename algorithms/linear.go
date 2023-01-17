package algorithms

import "fmt"

func LinearSort() {

	arrSize := 10
	hayStack := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	needle := 8

	fmt.Printf("Searching for %v\nDoes...\n", needle)
	for i := 0; i < arrSize; i++ {
		fmt.Printf("%v = %v? ", hayStack[i], needle)
		if hayStack[i] == needle {
			fmt.Printf("%v\n", "Yes!")
		} else {

			fmt.Printf("%v\n", "No!")
		}
	}
}
