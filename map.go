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

	fmt.Println("map的遍历")
	//map是hashmap，是无序的
	for k, v := range m {
		fmt.Printf("k:v = %s:%s\n", k, v)
	}
	for k := range m {
		fmt.Printf("k = %s\n", k)
	}

	name := m["name"]
	fmt.Println(name)

	//获取map中不存在的key，拿到的是zero value，即空字符串
	name1 := m["name1"]
	fmt.Println(name1)

	//判断key在不在
	name2, ok := m["name2"]
	fmt.Println(name2, ok) // false

	if name3, ok := m["name3"]; ok {
		fmt.Println(name3, ok) // name3不存在，所以不会被执行
	} else {
		fmt.Println("name3不存在")
	}

}
