package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "kevin",
		"age":  "29",
	}
	fmt.Println(m) //map[age:29 name:kevin]

	m2 := make(map[string]int) // m2 == empty map
	fmt.Println(m2)            //map[]

	//go语言的nil不像其他语言的nil，它是可以参与运算的，可以安全使用
	var m3 map[string]int // m3 == nil
	fmt.Println(m3)       //map[]

}
