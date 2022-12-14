package main

import "fmt"

var a = 11

// 自定义类型
type MyInt int32

func foo(n int) {
	a := 1
	a += n
}

func main() {
	// 变量声明
	// var a int = 1
	// b := 2
	foo(5)

	// var count = int(5)
	// for i := 0; i < count; i++ {
	// 	fmt.Println(i)
	// }

	// 自定义类型
	// var c = int32(1)
	// var d MyInt = c        // 虽然MyInt的底层类型是int32类型，但是本质上两个还是两种完全不同的类型，不能完全赋值，只能通过显示的类型转换
	// var e MyInt = MyInt(c) // 显示转换类型

	// 常量-const
	// 无类型常量 存在 隐式转换
	// 实现枚举
	// b，c没有赋值，所以隐式的重复上面a=1的赋值，所以都是1
	const (
		a = 1
		b
		c
	)

	const (
		_     = iota //_ 空白字符可以忽略，来实现枚举从1开始
		red          //1
		blue         //2
		black        //3
		_            // 4的位置用空白字符表示，就证明跳过了4
		white        //5
	)
	fmt.Println("=====常量=====", a, b, c, red, blue, black, white)

	// 数组
	var arr [4]int
	var strList = [6]string{"0", "1", "2", "3", "4", "5"} // 初始化值

	var lenList = [...]int{
		1, 2, 3, 4,
	} // 省略长度自动推算

	lenList2 := [4]int{
		1, 2, 3, 4,
	}

	var index = [...]int{
		13: 2,
	} // 通过下标赋值

	// var arr2 [2][3][4]int // 多维数组

	fmt.Println("=====数组=====", arr, strList, lenList)
	fmt.Println("=====数组=====", arr[0], len(arr))    // 获取数组长度
	fmt.Println("=====数组=====", index, len(index))   // 获取数组长度
	fmt.Println("=====数组=====", lenList == lenList2) // 多维数组

	// 切片
	var spliceArr []int

	var spliceArr2 = make([]string, 2)

	// 基于已有数组切片
	spliceArr3 := strList[1:3] // 开始下标为1，截取长度为3-1

	fmt.Println("=====切片=====", spliceArr2, spliceArr3)

	fmt.Println("=====切片=====", spliceArr, len(spliceArr), cap(spliceArr)) // 切片长度和容量
	spliceArr = append(spliceArr, 1)
	fmt.Println("=====切片=====", spliceArr, len(spliceArr), cap(spliceArr)) // 切片长度和容量

	// map
	var m map[string]int
	var m1 = map[string]int{
		"哈哈哈": 2,
	}
	m2 := map[[4]int]string{
		lenList: "哈哈哈",
	}
	fmt.Println("======map=====", m, m1, m2)

	for key, value := range m2 {
		fmt.Println(key, value)
	}

	// 自定义类型
	type T int
	type T2 T

	var t T = 1
	t1 := T2(t)
	t2 := "hello"
	// t3 := T2(t2) // cannot convert t2 (variable of type string) to type T2 底层类型不同。T2的底层类型是int
	fmt.Println("=====自定义类型 type=====", t1, t2)

	// 结构体 struct
	type Person struct {
		title string // 如果有重复字段，在嵌入式的情况下会优先访问父结构体的字段
		name  string
		age   int
		sex   string
	}

	type Book struct {
		title string
		page  int
		index map[string]int
		Person
	}
	book := Book{
		title: "哈哈哈",
		page:  1,
		Person: Person{
			title: "1",
			name:  "哈哈哈",
			age:   2,
			sex:   "男",
		},
	}

	book.index = map[string]int{
		book.title: book.page,
	}
	fmt.Println("====结构体====", book.title, book.Person, book)

	var s = []int{1, 2, 3, 4}

	for i, value := range s {
		fmt.Println(i, value)
	}
	h := "这是一个字符串"
	for i, v := range h {
		fmt.Println(i, v, string(v))
		// 0 36825 这
		// 3 26159 是
		// 6 19968 一
		// 9 20010 个
		// 12 23383 字
		// 15 31526 符
		// 18 20018 串
	}

	// for i := 0; i < len(h); i++ {
	// 	// Go 语言中的字符串值也是一个可空的字节序列，字节序列中的字节个数称为该字符串的长度。一个个的字节只是孤立数据，不表意。
	// 	fmt.Println(i, h[i])
	// }

outerLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break outerLoop
			}
			fmt.Println("====break=====", i, j)
		}
	}
}
