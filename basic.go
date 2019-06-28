package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

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
	euler()
	triangle()
}

func triangle(){
	var a,b int =3,4
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
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

func euler(){
	c:=cmplx.Pow(math.E,1i*math.Pi)+1
	fmt.Printf("%.3f",c)
}