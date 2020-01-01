package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

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

	euler()
	triangle()
}

/**
go的类型转换是强制的，没有隐式类型转换
*/
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) //float算出来有可能是4.99999999994,这样强制转换为int会是4，该怎么办
	fmt.Println(c)
}

/**
用go验证欧拉公式
复数的使用
*/
func euler() {
	//(0+1.2246467991473515e-16i),不是整数0的原因是complex64的实部和虚部都是float32，而complex128的实部和虚部都是float64
	//浮点数标准都是不准的
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)         //math.E是一个非常特殊的数字，这样的写法更好
	fmt.Printf("%0.3f\n", cmplx.Exp(1i*math.Pi)+1) //(0.000+0.000i)
}
