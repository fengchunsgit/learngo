package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 :=[3]int{3,5,7}
	arr3:=[...]int{2,4,5,6,8,10}
	var grid [4][5]int
	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)
}
