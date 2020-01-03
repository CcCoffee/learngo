package basic

import "fmt"

//算法：寻找最长的不含有重复字符的子串 abcabcbb -> abc， bbbbb->b, pwwkew -> wke
//对于每一个字母x
//* lastOccurred[x]不存在，或者 < start ->无需操作
//* lastOccurred[x] >= start -> 更新start
//* 更新lastOccurred[x]，更新maxLength

//主要时利用了map进行查字段
func findMaxLengthFromString(content string) int {
	//()(start)()()(*)()()
	//-->不断向前移动
	//()(start)()()()(*)()
	//for循环遍历每个字符，go没有char

	//start指针的移动由lastOccurred[x]是否存在决定
	start := 0
	lastOccurred := make(map[rune]int)
	maxLength := 0

	for i, ch := range []rune(content) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength { //maxLength只在出现更长的不重复字符子串时才进行更新
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(findMaxLengthFromString("abbcdee"))
	fmt.Println(findMaxLengthFromString("aa"))
	fmt.Println(findMaxLengthFromString(""))

	fmt.Println(findMaxLengthFromString("中中"))
	fmt.Println(findMaxLengthFromString("中人中"))
	fmt.Println(findMaxLengthFromString("中文名中心"))
}
