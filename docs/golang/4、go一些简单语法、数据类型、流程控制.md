# 基本数据类型

>- 程序实体和关键字
>- 变量和常量
>- 整数类型的命名和宽度,整数类型值得表示法
>- 浮点数类型
>- 复数类型
>- byte 与 rune
>- 字符串类型

>- 数组类型
>- 切片类型
>- 切片的更多操作方法
>- 字典类型
>- 通道类型
>- 通道的更多种类

>- 函数
>- 结构体和方法
>- 接口
>- 指针

>- if语句
>- switch语句
>- for语句
>- select语句

>- defer语句
>- 异常处理error
>- 异常处理panic
>- go语句

> go Playground:
在浏览器里面编辑并运行go语言代码。http://play.golang.org. 

## 变量
- 使用关键字`var`定义变量，自动初始化为零值。
- 如果提供初始化值，可省略变量类型，由编译器自动推断。

```go
var x int
var f float32 = 1.6
var s = "abc"
```

- 在函数内部，可用更简略的 `:=` 方式定义变量。

```go
func main() {
    x := 123
}
```

- 可以一次性定义多个变量, 函数形参也可以这样

```go
var x, y, z int
var s, n = "abc", 123

var (
    a int
    b float32
)

func main() {
    n, s := 0x1234, "Hello, World!"

}
```

- 编译器会将未使 的局部变量当做错误。

- 常量值必须是编译期可确定的数字、字符串、布尔值。

```go
const s = "Hello, World"
const (  //常量组
    a, b = 10, 100
    c bool = false
)
```
> 未使局部常量不会引发编译错误。
> 在常量组中，如不提供类型和初始化值，那么视作与上一常量相同。
> 常量值还可以是 len、cap、unsafe.Sizeof 等编译期可确定结果的函数返回值。
> 如果常量类型以存储初始化值，那么不会引发溢出错误。


- 枚举
 
- 关键字 iota 定义常量组中从 0 开始按计数的递增枚举值。

```go
const (
    Sunday = iota // 0
    Monday    //1，通常省略后续表达式。 
    Tuesday   //2
    Wednesday //3
    Thursday  //4
    Friday    //5
    Saturday  //6
)
```


## 流程控制

### if

- if的条件里不需要括号

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func fileNameIf() {
	const filename = "aaa.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println("Hello World")
}
```

- if的条件里可以赋值
- if的条件里赋值的变量作用域就在这个if语句里

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "aaa.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
```

### switch

- switch 会自动 break，除非使用 fallthrough

```go
func calculator(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupproted operator:" + op)
	}
	return result
}
```

- switch 后可以没有表达式

```go
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}
```


### 循环

- for 的条件里不需要括号

- for 的条件里可以省略初始条件，结束条件，递增表达式

```go
func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
}
```

```go
import (
	"fmt"
	"strconv"
	//	"strings"
	"math"
)

// 正整数转成二进制
func convertToBinary(n int) string {
	result := ""
	if n%2 != 0 {
		result = "0"
	}

	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	//高位补0,计算机内部表示数的字节单位是定长的，
	//如8位，16位，或32位，位数不够时，高位补零。

	if len(result)%8 != 0 {
		var mend int
		if len(result)/8 == 0 {
			mend = len(result) - 8
		} else {
			mend = len(result) - (len(result)/8)*8
		}

		for i := 0; i < int(math.Abs(float64(mend))); i++ {
			result = "0" + result
		}
	}
	return result
}
```


