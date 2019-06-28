package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func converToBin(n int) string{
	result:=""
	for;n>0;n/=2{
		lsb:=n%2
		result=strconv.Itoa(lsb)+result
	}
	return result
}

func main() {
	fmt.Println(
		converToBin(5),//101
		converToBin(13),// 1 0  1 1 -->1101
		converToBin(7238788500))
	printFile("abc.txt")
}

func printFile(filename string){
	file,err:=os.Open(filename)
	if err!=nil{
		print(err)
	}

	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}
