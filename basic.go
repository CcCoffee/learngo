package main

import "fmt"

//定义包内部的变量，go没有全局变量的说法
//var aa = 3
//var ss = "kkk"
//bb:=true //错误，函数的外面定义变量不可以使用:=
//===> 等效于
var ( //可以省略var
	aa = 3
	ss = "kkk"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Printf("%d %d %q\n", a, b, s)
}

/**
编译器可以自动决定类型
*/
func variableTypeDeduction() {
	var a, b, c = 3, true, "abc"
	fmt.Println(a, b, c)
}

/**
 * 推介
 * 使用:=进行定义和初始化变量，只能在函数内使用
 */
func variableShorter() {
	a, b, c := 3, true, "abc"
	fmt.Println(a, b, c)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss)
}
