package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	// 1. 创建slice
	var s []int           //定义一个slice s，go语言中每个变量一旦定义，就会分配zero value，对于slice类型的zero value是nil
	fmt.Println(s == nil) //true

	//打印奇数
	for i := 0; i < 100; i++ {
		//填满底层数组后，就会扩充*2
		//len=63, cap=64
		//len=64, cap=64
		//len=65, cap=128
		//len=66, cap=128
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	//使用make生成元素为0的slice
	s2 := make([]int, 16) //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	printSlice(s2)

	//16为初始长度，32为capacity
	s3 := make([]int, 16, 32) //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=32
	printSlice(s3)
}
