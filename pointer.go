package main

import "fmt"

/**
交换两个值
*/
func swap(a, b int) {
	b, a = a, b //go中只有值传递没有引用传递，没用
}

/**
go的指针*放在类型前面
*/
func swap1(a, b *int) {
	*b, *a = *a, *b //go中只有值传递没有引用传递，没用
}

//更好的方法
func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	swap(a, b)
	println(a, b) //3 4,无变化
	swap1(&a, &b)
	println(a, b) //4 3
	c, d := 5, 6
	q, r := swap2(c, d) //6 5
	fmt.Println(q, r)
}
