# interface

- 接口（interface）无法被实例化.

- 接口类型也算是引用类型。

- 接口是根据方法集合区分的（定义的所有方法）。

- 代表某一类特征（`Duck Typing`）

## 接口的使用

```
type Animal interface {
    Eat(foot string)
    Run()
}
```

```
type Person struct {
    
}
func (p *Person)Eat(foot string) {
    
}
func (p *Person)Run() {
    
}
```

- Person 实现了 `Eat(foot string)` 方法和 `Run()` 方法, 也就是说 Person 实现了接口 Animal ，并且这种实现接口的方式是无侵入的.

- 判断一个类型是否实现了某个接口：
    - 两个方法的签名需要完全一致
    - 两个方法的名称要一模一样


## 接口的另一种使用

- 可以把实现接口的类型，赋值给对应的接口类型变量

```
var p Person = &Person{}
var a Animal = p
```

- 变量 `p` 是接口变量 `a` 的`实际值`，也叫`动态值`.
- 变量 `p` 的类型 `Person` 叫做接口变量 `a` 的`实际类型`，也叫 `动态类型`.

- 类型 `Animal` 是接口变量 `a` 的 `静态类型`，是永远不会变化的。（相对于动态类型是变化的跟我们的赋值有关）

- 接口类型变量的零值是`nil`。

- 当给一个接口变量赋值的时候，该变量的动态类型和动态值会一起被存储在一个专用的数据结构中。
- 其实，这个接口变量的值是这个专用数据结构的一个实例，而不是我们赋给改变量的那个实际值。

- 这个专用的数据结构在 Go 语言的 `runtime` 包中叫 `iface` 。`iface`的实例会包含两个指针，一个是指向类型信息的指针，另一个是指向动态值的指针。这里的类型信息是由另一个专用数据结构的实例承载的，其中包含了动态值的类型，以及使它实现了接口的方法和调用他们的途径。


- 使用 `==` 判断接口变量是否为 `nil`：

```
var p Person
var a Animal = p
// p == nil , a != nil
if a == nil {
	fmt.Println("a == nil")
}else{
	fmt.Println("a != nil")
}
```

- 对一个接口变量只声明不初始化，或者直接给接口变量赋值 `nil`，这时接口变量值为 `nil`。


## 接口的组合

接口的组合：接口类型之间的嵌入。

建议声明小接口，更容易组合接口，扩展性强、比较灵活。

组合的接口之间有同名的方法（方法签名不同也不行）就会编译报错。


看Go标准库 `io` 包：

```
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}
```
- 同时实现了 `Reader` 接口和 `Writer` 接口，就相当于实现了组合接口 `ReadWriter`。


```
type Closer interface {
	Close() error
}
type ReadCloser interface {
	Reader
	Closer
}
type WriteCloser interface {
	Writer
	Closer
}
```


## nil 接口变量

```
var p Person = nil
var a Animal = p
// p == nil , a != nil
```

Person 实现了接口，p 是一个nil变量，赋值给接口类型变量 a，a 依然可以调用接口实现的方法。方法接收者必须是指针类型才能调用接口实现的方法。


## Duck Typing

`Duck Typing` ，看起来像鸭子，它就是鸭子。

- `Go`语言对象没有继承和多态，只有封装性。

- 协议注重关注方法的实现，而很少关注类型。

- 只要实现了B协议方法的实体，是B协议类型的入参， 都可以作为的传入。

- 变相实现了继承和多态，实现多个协议，就可以是类似多继承。


