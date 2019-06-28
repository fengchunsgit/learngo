package main

import "fmt"

func div(a,b int) (q,r int){
	return a/b,a%b
}
func main() {
	fmt.Println(div(13,4))
	q, r :=div(13,3)
	fmt.Println(q,r)
	fmt.Println(sum(1,2,3,4,5,5,6))
}

func sum(numbers ...int) int{
	s:=0
	for i:=range numbers {
		s+=numbers[i]

	}
	return s
}
