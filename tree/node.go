package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//go中没有构造方法，如果像加以控制，可以使用工厂函数代替
//工厂函数一般返回结构的地址，结构不需要考虑在哪里分配，只要局部变量产生一个这样的结构，再把地址返回就可以
//结构创建在堆上还是栈上？
//在c++中局部变量是在栈上的，函数一但退出局部变量就会立刻被销毁，如果要传出去必须在堆上分配。堆上分配的变量需要手动释放。
// java几乎所有东西都是分配在堆上的，都要用new，分配在堆上才会有垃圾回收机制。
// 使用go语言不需要知道结构创建在堆上还是栈上。当编译器看到返回变量要给别人使用，则这个treeNode会在堆上去分配。如果没有则会在栈上分配。
func createNode(value int) *treeNode {
	return &treeNode{value: value} //这里返回了一个局部变量的地址给外面用，在c++中这是不可以的。go中局部变量也可以返回给外部使用
}

func main() {
	//go语言的构造方法
	var root treeNode
	fmt.Println(root) //{0 <nil> <nil>}
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) //c中right为指针，所以必须root.right->left操作。go则不管是指针还是实例都用`.`访问成员
	root.left.right = createNode(2)
	nodes := []treeNode{
		{value: 3},
		{},
		{5, nil, &root},
	}
	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {5 <nil> 0xc0000a0000}]
}
