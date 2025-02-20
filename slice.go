package slice

import "fmt"

func PrintSlice(slice []int) {
	for _, val := range slice {
		fmt.Printf("%d ", val)
	}
}

func PrintSliceString(slice []interface{}) {
	for _, val := range slice {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func PrintSlice2(slice [][]int) {
	for _, rows := range slice {
		for _, val := range rows {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
