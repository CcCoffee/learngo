package container

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

	// 2. 复制slice
	copy(s2, s1)
	printSlice(s2) //[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16

	// 3. 删除slice的元素
	// 删除下标为3的元素8，s2[4:]...使用了go中的展开语法
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2) //[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0], len=15, cap=16

	// 4. 删除slice的头元素
	s2 = s2[1:]
	printSlice(s2) //[4 6 0 0 0 0 0 0 0 0 0 0 0 0], len=14, cap=15

	// 5. 删除slice的尾元素
	printSlice(s1) //[2 4 6 8], len=4, cap=4
	s1 = s1[:(len(s1) - 1)]
	printSlice(s1) //[2 4 6], len=3, cap=4
}
