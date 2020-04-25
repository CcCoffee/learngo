## 软件开发的新挑战

1. 多核硬件架构
2. 超大规模分布式计算集群
3. Web 模式导致的前所未有的开发规模和更新速度

## Go 语言简介

Go 语言的雏形诞生于 2007 年的谷歌，go 语言的发明是为了解决谷歌内部面前的以上三个问题。出现了开发的效率与大规模系统开发的协作的问题。

### Go 语言的创始人

Rob Pike：Unix 的早期开发者，UTF-8 创始人

Ken Thompson：Unix 的创始人、C 语言创始人、1983 年获图灵奖

Robert Griesemer：Google V8 JS Engine 和 Hot Spot 开发者

> Go 语言是出自大师之手的杰作。

### 语言的特点

* 简单：C 有 37 个关键字，C++ 11 有 84 个关键字，Go 目前只有 25 个关键字。而且 Go 的创始人一直坚持 Go 语言会只有这些关键字。

* 高效：Go 是编译的强类型语言，和 Java 不同的是 Go 在支持了垃圾回收的同时为了提供更高效的内存访问，特别提供了使用指针进行直接的内存访问。

* 生产力：Go 不仅语法简介还有特别的接口类型，还有一些编程约束直接就为开发者做出一些更好的选择。

  例如在程序的拓展上，一般的语言会支持复合和继承，很多的面向对象编程的语言的书都会谈论复合大于继承，然后解释这是为什么。在 Go 语言当中直接就只支持了复合。越来越多的应用都采用 Go 语言来开发，常见的就有 docker、kubernetes。由于云端大量的使用了 kubernetes 和 docker，所以 go 语言也被称为了云计算语言。区块链是最近继 AI 以来最热门的话题了，非常热门的以太坊、HyperLedger 都是可以采用 Go 语言来开发的，所以 Go 语言也被称为区块链开发语言。

## 准备开始 Go 冒险之旅

### 下载安装 Go 语言

https://golang.org/doc/install

Https://golang.google.cn/dl/

### 安装 IDE

Atom: https://atom.io + Package: go-plus

## 编写第一个 Go 程序

### 开发环境构建

#### GOPATH

1. 在 1.8 版本前必须设置这个环境变量

2. 1.8 版本之后（含 1.8）如果没有设置使用默认值

   在 Unix 上默认为`$HOME/go`，在 Windows 上默认为`%USERPROFILE%/go`

   在 Mac 上 GOPATH 可以通过 `~/.bash_profile`来设置

> 使用`go version` 查看 go 版本



go_learning/src/ch1/main/hello_world.go

```go
package main // 包，表示代码所在的模块（包）

import "fmt" // 引入代码依赖

// 功能实现
func main() {
	fmt.Println("Hello world")
}
```

使用 `go run` 直接运行 go 文件

```bash
$ go run hello_world.go 
Hello world
```

使用 `go build` 编译文件

```bash
$ go build hello_world.go
$ tree .
.
├── hello_world
└── hello_world.go
```

go 默认情况下都会选用静态连接，所以说编译完的程序都会是一个独立的二进制文件，有非常好的便携性，可以复制到不同的机器上运行。在安装部署的时候，尤其是通过容器去安装部署的时候便携能力是非常强的。

### 应用程序入口

1. 必须是 main 包：`package main`

   go 语言里 package 名字跟目录的名字不需要保持一致，文件夹名不一定要是 main。

2. 必须是 main 方法：`func main()`

3. 文件名不一定是 main.go

### 程序退出状态

一般情况下 main 方法的返回值就可反映程序退出的状态。在 Go 语言中 main 函数是不支持返回值的，我们需要通过一个 `os.Exit` 来返回状态。

```go
package main // 包，表示代码所在的模块（包）

import (
	"os"
	"fmt" // 引入代码依赖
)

// 功能实现
func main() {
	fmt.Println("Hello world")
	os.Exit(1)
}
```

使用 `go run` 直接运行 go 文件

```bash
$ go run hello_world.go
Hello world
exit status 1
```

### 获取命令行参数

在程序中直接通过 `os.Args` 获取命令行参数

```go
func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello world", os.Args[1])
	}
}
```

使用 `go run` 直接运行 go 文件

```bash
$ go run hello_world.go 2
Hello world 2
```

