package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//go中没有构造方法，如果像加以控制，可以使用工厂函数代替
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
	root.right.right = createNode(10)
	nodes := []treeNode{
		{value: 3},
		{},
		{5, nil, &root},
	}
	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {5 <nil> 0xc0000a0000}]
}
