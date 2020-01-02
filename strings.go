package main

import "fmt"

func main() {
	//参考 https://blog.csdn.net/zhusongziye/article/details/84261211
	fmt.Printf("%U\n", '我') //unicode编号为U+6211 对应10进制数为 25105, 对应二进制数为0110 0010 0001 0001，编号本身需要2个字节表示
	//25105落在utf8的三字节范围0x800-0xffff中，即utf8使用三字节编码。对应的编码格式为1110XXXX 10XXXXXX 10XXXXXX
	//使用unicode编号的二进制表示填充编码格式的'X'部分
	//	10XXXXXX -> 1001 0001 剩余0110 0010 00XX XXXX
	//	10XXXXXX -> 1000 1000 剩余0110 XXXX XXXX XXXX
	//	1110XXXX -> 1110 0110
	//=> 1110 0110 1000 1000 1001 0001 这就是'我'的UTF-8编码，占用了三个字节
}
