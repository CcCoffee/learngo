package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//与func print(node treeNode)能达到差不多的效果，区别在与这个方法直接拓展了结构体，调用方式更友好
func (node treeNode) print() {
	fmt.Println(node.value)
}

//go的方法是值传递的，这个node对象已经是一个新对象，下面试图修改原本对象的属性无法实现
func (node treeNode) setValue1(value int) {
	node.value = value
}

//go语言没有this和self，如果要定义方法就必须给方法接收者取个名字
//使用指针作为方法接收者，只有使用指针才能改变结构内容
//nil指针也可以调用方法
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil node. ignore")
		return
	}
	node.value = value //go友好的地方在于，指针对象也可以直接用`.`访问对象属性
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
	root.print()       //3
	root.setValue1(33)
	root.print()      //3, 没有变成33，因为go的方法是值传递的
	root.setValue(33) //setValue方法需要传入的指针对象，go会自动把root转为指针对象传入，这里不需要用&root.setValue(33)来调用
	root.print()      //33

	//如果调用对象是指针对象，当go编译器知道调用的方法需要值，会根据指针对象从地址把值对象取出来；如果调用方法需要指针，就不用去值对象，直接给方法
	pRoot := &treeNode{}
	pRoot.setValue(100)
	pRoot.print() //100

	var pNode *treeNode
	pNode.setValue(200)
	pNode = &treeNode{}
	pNode.setValue(200)
	pNode.print()

}
