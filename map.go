package main

import "fmt"

func main() {
	m:=map[string]string{
		"name":"cc",
		"sex":"male",
	}
	m2:=make(map[string]int)  //m2==empty map

	var m3 map[string]int //m3==nil
	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("traversing map")
	for k,v:=range m{
		fmt.Println(k,v)
	}

	fmt.Println("getting value")
	name,ok:=m["name"]
	fmt.Println(name,ok)
}
