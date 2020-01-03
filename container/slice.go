package container

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

	fmt.Println("Extending slice")
	testArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	a1 := testArray[2:6] //[2 3 4 5]
	//a1为[2 3 4 5]，但是a2这里竟然可以获取到6。如果直接用a1[5]是无法获取的，即s[i]不可以超越len[s]。
	// 因为slice底层是view而已，只要不超过数组实际的cap(s)都可以拓展
	a2 := a1[3:5]                                                       //[5 6]，slice可以向后拓展，不可向前拓展。
	fmt.Printf("a1=%v, len(a1)=%d, cap(a1)=%d\n", a1, len(a1), cap(a1)) //a1=[2 3 4 5], len(a1)=4, cap(a1)=6 //cap(a1)=6,是因为有[2 3 4 5 6 7]
	fmt.Printf("a2=%v, len(a2)=%d, cap(a2)=%d\n", a2, len(a2), cap(a2)) //a2=[5 6], len(a2)=2, cap(a2)=3 //cap(a1)=3,是因为有[5 6 7]
	fmt.Println(a1[3:6])                                                //[5 6 7]
	//fmt.Println(a2[3:7])//panic: runtime error: slice bounds out of range

	fmt.Println("--------- 向slice添加元素 ---------")
	fmt.Println("testArray = ", testArray) //testArray =  [0 1 2 3 4 5 6 7]
	a3 := append(a2, 10)                   //                      a2 =            [5 6]
	fmt.Println("a3 = ", a3)               //                     a3 =            [5 6 10]
	a4 := append(a3, 11)
	fmt.Println("a4 = ", a4) //                     a4 =            [5 6 10 11]
	a5 := append(a4, 12)
	fmt.Println("a5 = ", a5) //                     a5 =            [5 6 10 11 12]
	//s4 and s5 no longer view testArray, system create a new longer array.
	//添加元素时如果超越cap，系统会重新分配更大的底层数组。原来的数组如果有人使用它还可以存在，否则等待垃圾回收
	//由于值传递的关系，必须接受append的返回值，s= append(s,10)
	fmt.Println("testArray = ", testArray) //testArray =  [0 1 2 3 4 5 6 10]
}
