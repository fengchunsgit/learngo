package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename="abc.txt"
	//if 的条件里可以赋值
	contents,err:=ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents)
	}
	fmt.Println(grade(0),grade(20),grade(100))
}


//switch 会自动break，除非使用fallthrough
func eval(a,b int,op string) int{
	var result int
	switch op{
	case "+":
		result=a+b
	case "-":
		result=a-b
	case "*":
		result=a*b
	case "/":
		result=a/b
	default:
		panic("unsupported operator:"+op)
	}
	return result
}

func grade(score int) string{
	g:=""
	switch {
	case score<60:
		g="f"
	case score<80:
		g="c"
	case score<90:
		g="b"
	case score<=100:
		g="a"
	default:
		panic(fmt.Sprintf("Wrong score:%d",score))
	}
	return g
}
