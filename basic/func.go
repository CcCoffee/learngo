package basic

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		//return a/b
		q, _ := div(a, b) //使用_代替不想使用的值，避免报错
		return q
	default:
		panic("unsupported operation : " + op)
	}
}

/**
直接使用panic不够优雅，使用error多值传递代替,程序不会中断
*/
func eval2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a/b
		q, _ := div(a, b) //使用_代替不想使用的值，避免报错
		return q, nil
	default:
		//panic("unsupported operation : " + op)
		return 0, fmt.Errorf("unsupported operation : %s\n", op)
	}
}

//返回多值
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//与div方法等效
//函数体长的时候看得不够清楚，不建议使用
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

/**
函数式编程
*/
func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name() //获取方法名
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

/**
可变参数列表
go没有默认参数，可选参数，函数重载。。。
*/
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		fmt.Println(numbers[i])
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(div(13, 3)) // 4 1
	q, r := div(13, 3)      //div(13, 3).var编辑器自动生成返回值名q, r := div(13, 3)
	fmt.Println(q, r)       // 4 1
	i, r2 := div2(11, 2)
	fmt.Println(i, r2) // 5 1
	if result, e := eval2(1, 1, "&"); e != nil {
		fmt.Println("Error : ", e)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(Pow, 3, 4)) //Calling function main.Pow with args (3, 4)
	//使用匿名函数，函数名自动取main.main.func1
	//go没有lambda表达式
	fmt.Println(apply( //Calling function main.main.func1 with args (3, 4)
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Printf("Sum = %d\n", sum(1, 2, 3))
}
