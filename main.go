package main

import "fmt"

var a = 11

// 自定义类型
type MyInt int32

func foo(n int) {
	a := 1
	a += n
	fmt.Println("a==", a)
}

func main() {
	// 变量声明
	// var a int = 1
	// b := 2
	fmt.Print(a)
	foo(5)

	fmt.Println("result a=", a)
	// var count = int(5)
	// for i := 0; i < count; i++ {
	// 	fmt.Println(i)
	// }

	// 自定义类型
	// var c = int32(1)
	// var d MyInt = c        // 虽然MyInt的底层类型是int32类型，但是本质上两个还是两种完全不同的类型，不能完全赋值，只能通过显示的类型转换
	// var e MyInt = MyInt(c) // 显示转换类型

}
