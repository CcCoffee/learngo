package main

import (
	"fmt"
	"unicode/utf8"
)

//编码：unicode编号转utf-8编码
//解码：utf-8编码转unicode编号
func main() {
	//参考 https://blog.csdn.net/zhusongziye/article/details/84261211
	fmt.Printf("%U\n", '我') //unicode编号为U+6211 对应10进制数为 25105, 对应二进制数为0110 0010 0001 0001，编号本身需要2个字节表示
	//25105落在utf8的三字节范围0x800-0xffff中，即utf8使用三字节编码。对应的编码格式为1110XXXX 10XXXXXX 10XXXXXX
	//使用unicode编号的二进制表示填充编码格式的'X'部分
	//	10XXXXXX -> 1001 0001 剩余0110 0010 00XX XXXX
	//	10XXXXXX -> 1000 1000 剩余0110 XXXX XXXX XXXX
	//	1110XXXX -> 1110 0110
	//=> 1110 0110 1000 1000 1001 0001 这就是'我'的UTF-8编码，占用了三个字节

	s := "Aa我是Kevin"
	fmt.Println(len(s)) //13,这里的s变量数据内容其实是经utf-8编码过的
	// len函数可以打印出字符串占用的字节数，utf-8编码的unicode中每个中文通常占三个字节byte，少数占用4个字节
	for _, ch := range []byte(s) {
		fmt.Printf("%X ", ch) //41 61 E6 88 91 E6 98 AF 4B 65 76 69 6E
	}
	fmt.Println()

	//rune 等同于int32，即4个字节长度,常用来处理unicode编号或utf-8字符，相当于go的char
	for i, ch := range s { //range把string进行了utf-8的解码，解出来每个字符又进行了转译为unicode编号，转完放在了rune这个四字节类型的变量ch里面
		// 注意： unicode编号是世界上所有字符的唯一编号，它仅占2个字节，但是转化为能够存储的utf-8编码格式时占用的字符数却不一定是2个字符
		// utf-8有可能占用1、2、3或4个字符，如ascii编码的字符仅占用1个字符。而中文大部分需要占用3个字符。具体编码过程参见上面的url链接文章
		fmt.Printf("(%d %X) ", i, ch) //(0 41) (1 61) (2 6211) (5 662F) (8 4B) (9 65) (10 76) (11 69) (12 6E)
	}
	fmt.Println()

	// 上面输出的是rune类型变量的16进制表示，这里使用%c直接输出对应的字符
	// 这里的下标i指的是字符对应字节数组的起始位置，它是不连续的
	for i, ch := range s {
		fmt.Printf("(%d %c) ", i, ch) //(0 A) (1 a) (2 我) (5 是) (8 K) (9 e) (10 v) (11 i) (12 n)
	}
	fmt.Println()
	// 直接输出第几个字符是谁。显示转换s字符串为rune的数组，注意下标的不同
	// 强转为rune数组，类型转换并不是对字节内容的重新理解，而是真正去转为decode的结果，
	// 这时候每个字符用的rune变量占用了4个字节，重新创建了rune数组存储，已经不是同一块内存
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch) //(0 A) (1 a) (2 我) (3 是) (4 K) (5 e) (6 v) (7 i) (8 n)
	}
	fmt.Println()
	//使用len只是获取字节长度，这里使用utf8.RuneCountInString(s)获取字符数
	fmt.Println("utf8.RuneCountInString(s)打印字符数：", utf8.RuneCountInString(s)) //9

	// strings的使用
	// fields,split,join
	// contains, index
	// toLower, toUpper
	// trim, trimRight, trimLeft

}
