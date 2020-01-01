package main

import "fmt"

//go中的数组是值类型，会复制数组。所以我们必须定义数组的长度。
// []int 与 [4]int在go中是完全不一样的东西
func max(numbers [4]int) (int, int) {
	maxi := -1
	maxValue := 01
	for i, v := range numbers {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	return maxi, maxValue
}

func main() {
	//数量写在类型前面
	var arr1 [5]int
	arr2 := [3]int{1, 3, 4}       //:=需要给初始值
	arr3 := [...]int{2, 3, 4, 5}  //让编译器去数总数
	fmt.Println(arr1, arr2, arr3) //[0 0 0 0 0] [1 3 4] [2 3 4 5]

	//二维数组
	var grid [4][5]int
	fmt.Println(grid) //[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]

	//遍历数组方式一
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	//遍历数组方式二
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	//遍历数组方式三
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//遍历数组方式四, 下划线省略变量
	for _, v := range arr3 {
		fmt.Println(v)
	}

	maxi, maxValue := max(arr3)
	fmt.Println(maxi, maxValue)

	//max(arr1)//./arrays.go:51:5: cannot use arr1 (type [5]int) as type [4]int in argument to max

}
