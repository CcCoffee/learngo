package main

import "fmt"

//go语言一般不用数组而用切片
//slice本身没有数据，是对底层array的一个view
func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println("arr[2:4] =", arr[2:4]) //arr[2:4] = [3 4]
	fmt.Println("arr[2:] =", arr[2:])   //arr[2:] = [3 4 5]
	fmt.Println("arr[:4] =", arr[:4])   //arr[:4] = [1 2 3 4]
	fmt.Println("arr[:] =", arr[:])     //arr[:] = [1 2 3 4 5]

	s1 := arr[2:]
	fmt.Println("arr[2:] =", s1)
	updateSlice(s1)
	fmt.Println("After updateSlice(s1)")
	fmt.Println(s1)  //[100 4 5]
	fmt.Println(arr) //[1 2 100 4 5]

	fmt.Println("Reslice")
	fmt.Println(s1)
	s1 = s1[2:]
	fmt.Println(s1)
}
