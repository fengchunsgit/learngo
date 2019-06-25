package main

import "fmt"

var (
	aa=11
	bb=333
	cc="sss"
)
func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,bb,cc)
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

func variableTypeDeduction(){
	var a,b,c,s=3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func variableShorter(){
	a,b,c,s:=3,4,true,"def"
	fmt.Println(a,b,c,s)
}