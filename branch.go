package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	//test1(filename)
	test2(filename)
	fmt.Println(grade(30))
	fmt.Println(grade(111))

}

/**
switch后面可以没有表达式，在case里面加入条件就可以
*/
func grade(score int) string {
	result := ""
	switch {
	case score < 60:
		result = "F"
	case score < 80:
		result = "C"
	case score < 90:
		result = "B"
	case score < 100:
		result = "A"
	default:
		//使用panic会中断程序执行，报错
		panic(fmt.Sprintf("Wrong score: %d\n", score))
	}
	return result
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
