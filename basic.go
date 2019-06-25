package main

import "fmt"

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
}

func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}

func variableInitialValue(){
	var  a,b int=3,4
	var s string ="abc"
	fmt.Println(a,b,s)
}