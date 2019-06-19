## Go语法注意点

- #### 1. 如果提供初始化值，可省略变量类型，由编译器自动推断。

```go
//定义变量 a
a := "word"
```

如果前面已经定义了变量，后面就不能使用这种方式定义变量

```go
//定义变量 a
var a string

//这样编译器会报错
a := "word"  

//应该这样
a = "word" 
```

- #### 2. 一句语句结束不用使用分号，使用换行就行；如果两句在一行，两句之间要使用分号。

```go
# 一般条件语句中使用这种的较多
a := 2; a < 3; a++ 
```

- #### 3. 变量定义，变量名在前，类型在后。

```go
var a int
```

- #### 4. 可以一次定义多个变量。

```go
var a, b, c int

//很常用的方式
var (
    d int
    e string
    f float32
)
```

- #### 5. 编译器会将未使用的局部变量当做错误。

- #### 6. 多个返回值，其中一个返回值不需要，可以使用 `_` 忽略。

```go
//函数 logInfo 返回两个值
func logInfo() (string, error) {
    
}

//调用 logInfo 函数, 不需要第一个返回值，需要第二个返回值，就可以使用 _ 忽略掉第一个返回值
_, err := logInfo()
```

- #### 7. 默认首字母大写是公开，小写是私有的。

```go
//函数 logInfo 就是私有的，只有包内部可以调用
func logInfo() (string, error) {
    
}

//函数 LogError 就是公开的，包外部文件可以调用
func LogError() (string, error) {
    
}

//属性也是相同, User 就是公开的，属性Name也是公开的
type User struct {
    Name string
}

//admin 就是私有的
type admin struct {
}
```

- #### 8. switch 语句，会自动 break.

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
- #### 9. for 循环条件里不需要括号.

```go
for i := 1; i <= 100; i++ {
	sum += i
}
```

- #### 10. for 条件可以省略初始条件。没有while；

```go
func convert(n int) {
    for ; n > 0; n /= 2 {
	}
}

for {

}
```

- #### 11. 可以使用 `+` 号直接对字符串拼接.

```go
a := "hello"
b := a + "word!"
```

- #### 12. 在不同类型之间赋值时，需要显式转换。 

```go
var x int = 3
var y int64
//必须显示类型转换
y = int64(x)
```

- #### 13. 没有条件的 switch

同 switch true 一样.
与编写长的 if-then-else 链比，这种形式更清晰。

```go
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```
