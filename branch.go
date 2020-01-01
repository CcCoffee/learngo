package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	//test1(filename)
	test2(filename)
}

/**
if的条件是可以赋值的
条件中赋值的变量作用域就在这个if语句里
*/
func test2(filename string) {
	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Printf("%s\n", contents)
	} else {
		fmt.Printf("%s %s\n", "file not exists", err)
	}
	//fmt.Println(contents) //编译器异常
}

func test1(filename string) {
	contents, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Printf("%s\n", contents)
	} else {
		fmt.Printf("%s %s\n", "file not exists", err)
	}
}
