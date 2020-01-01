package main

//import也可以用()
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	//省略初始条件
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	//省略递增条件
	//相当于while，go没有while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/**
死循环
go语言并发编程go routine都是死循环
*/
func forever() {
	for {
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1011 -> 1101
	)
	printFile("abc.txt")
	//forever() //while(true)
}
