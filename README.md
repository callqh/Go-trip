# Go 笔记

## Go 模块的进程

- 无法锁定依赖,就导致没办法进行可重现的构建（可重现的构建：因为你在本地开发，和别人拉去代码后在另一个电脑上开发时依赖的包可能会变化，所以可能没办法重现问题）
- `vendor` 目录,将依赖缓存到该目录，需要上传到 `git` 仓库用来锁定依赖，繁琐让项目变大
- `Go Module`,采用 `go.mod` 和 `go.sum` 来记录依赖（类似 `package.json`)**使用最小版本机制来进行包依赖选择**
  - `go mod init` 初始化 `Go Module` 模式，创建 `go.mod` 文件 -` go mod tidy` 自动记录项目中用到的依赖，输入到 `go.mod` 文件中
  - `go get xxxx` 为当前项目新增依赖； `xxx@version` 可以新增指定版本的依赖

## 数据类型

### 基本数据类型

- 数值类型
- 字符串类型

### 复合类型

#### 数组

`var [N]T` 长度固定(`N`)、类型相同(`T`)

- 数组类型不仅是逻辑上的连续序列，而且在实际内存分配时也占据着一整块内存
- Go 中传递数组的方式都是**纯粹的值拷贝**（Go 中的数组代表的就是整个数组，而不是指针），这会带来额外的内存拷贝开销，如果想通过指针的方式传递，那么就使用**切片**

#### 切片

`var []T`

- 没有固定长度，可以通过`nums = append(nums,1)`来动态添加元素
- 对已有数组进行切片`splice = arr[low,hight,max]` `low`是其实下标，切片长度为`(hight-low)`，max 是切片容量（常省略 max），对切片的操作会影响原数组
- 在基于数组进行切片后，当使用 append 追加变量超过原数组的边界时，切片会与原数组进行解绑，即之后切片的操作不会影响到原数组。这是因为 append 在自动扩容的时候，发现原数组容量已经不符合，就会新创建一个数组区域，将原数组的元素复制到新的区域中，因此进行解绑

#### `map`

`var m map[key_type]value_type`

- `map`的`key`的类型有所限制，仅限于可以使用`==`和`!=`进行比较的类型。**函数、切片以及`map`本身不能作为`key`值**
- `map`类型声明后必须赋予初值，后面才可以操作，否则程序会`panic`

**相关操作：**

- 新增： `map[key]=value`， 若果`key`已经存在，则会覆盖原来的值
- 获取： `map[key]`
- 查找： 在 `Go` 语言中，请使用“`comma ok`”惯用法对 `map` 进行键查找和键值读取操作。

```go
  m := make(map[string]int)
  value,ok := m["key"] // 如果这里不关心value值，可以用空值代替_
  if !ok {
    // 如果不存在
  }
```

- 删除： `delete(map,key)` 即时`key`不存在`delete`函数仍会正常执行
- 遍历：

```go
  for key, value := range m2 {
		fmt.Println(key, value)
	}
```

#### `struct`

```go
  type T struct {
    name string
    age int
  }
```

- 嵌入字段`（Embedded Field）`： 在访问的时候可以直接省略类型名，直接访问嵌入类型的元素`U.name`或者`U.T.name`

```go
  type T struct {
    name string
    age int
  }
  type U struct {
    T // 嵌入类型
    extra string
  }
```

- 结构体类型值初始化：
  ```go
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
  ```

## 变量

`var 变量名称 变量类型 = 值`

`var a int = 1`
`var a,b,c = 1,2,3`

- 变量块声明

```go
  var (
    a int = 1
    b string = "hello"
  )
```

- 短变量声明： `a:=1` (适用局部变量)

## 常量

常量的类型只局限于基本数据类型：数值、字符串和布尔类型。

- `untyped constant`(无类型常量): 变量的类型是不确定的，可以是任意类型的值，存在隐式转换

- 常量实现枚举
  - 隐式重复上一行
  - `iota` 行偏移量（类似数组的指针）

```go
    // b，c 没有赋值，所以隐式的重复上面 a=1 的赋值，所以都是 1
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
```

## 控制结构

### 分支

- 普通`if`

```go
	if bool {

	} else if xxx {

	} else {

	}
```

- 在 `if`中声明局部变量

```go
  if a := true; a {
    fmt.Println(a)
  }
```

- 快乐路径(`happy path`)

```go
func doSomething() error {
	if errorCondition1 {
		// 错误的逻辑处理 ... ...
		return err1
	}
	// 正确的逻辑处理 ... ...
	if errorCondition2 {
		//  错误的逻辑处理 ... ...
		return err2
	}
	// 正确的逻辑处理  ... ...
	return nil
}
```

> `Go` 社区把这种 `if` 语句的使用方式称为 `if` 语句的“快乐路径`（Happy Path）`”原则，所谓“快乐路径”也就是成功逻辑的代码执行路径，它的特点是这样的：
>
> - 仅使用单分支控制结构；当布尔表达式求值为 `false` 时，也就是出现错误时，在单分支中快速返回；
> - 正常逻辑在代码布局上始终“靠左”，这样读者可以从上到下一眼看到该函数正常逻辑的全貌；（靠左是指正确的逻辑代码不放在 if 里面）
> - 函数执行>到最后一行代表一种成功状态。

### 循环

`Go`中仅提供了一种循环那就是`for`。

- 普通 `for` 循环
  ```go
    for i:=1; i<10;i++{
      fmt.Println(i)
    }
  ```
- 始终为 `true` 的循环
  ```go
    for {
      // 循环体
    }
  ```
- `for range`
  用来对**复合结构**进行遍历（原生的 `string` 类型也可以)

  ```go
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
  ```

  - 在遍历字符串时，`v `获取的是 `Unicode` 码值，而不是字符本身。
  - `i` 为该 `Unicode` 字符码点的内存编码（`UTF-8`）的第一个字节在字符串内存序列中的位置


- 中断循环
  - `continue`：支持带`label`的`continue`
    ```go 
      outerLoop:
        for i := 0; i < 10; i++ {
          for j := 0; j < 10; j++ {
            if j == 2 {
              continue outerLoop // 直接中断内层j循环，跳转到外层i继续遍历
            }
            fmt.Println(i, j)
          }
        }
    ```
  - `break`：同样支持`label`
