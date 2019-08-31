# 读 "优雅关闭的 Go Web 服务器"

> 文章原文链接：[Go语言中文网](https://mp.weixin.qq.com/s/rA_oh472ZhfcAsAkWyyXFA)

文章《优雅的关闭 `Go Web` 服务器》写到：可以通过开启一个单独的 `goroutine` 拦截关闭信号，这样，当服务器真正关闭之前，可以做一些任务，做完任务再发出执行完毕的关闭号。

一些任务比如：清理某些资源；完成数据库事务；一些其他长时间的操作；退出服务的那一刻，刚好收到一个响应，为了保证所有请求完成，就可以在这里，在最大超时时间内去处理这个响应；dump进程当前状态；记录日志的动作。

启动应用，`Ctrl + C` 中断程序，中断信号被拦截，do something.....

```bash
go run main.go -listen-addr=5001
http: 2019/08/31 11:34:30 Server is ready to handle requests at :5001
^Chttp: 2019/08/31 11:34:32 Server is shutting down...
do something start .....  2019-08-31 11:34:32.594668 +0800 CST m=+2.337451148
do something end .....  2019-08-31 11:34:37.598248 +0800 CST m=+7.340881516
http: 2019/08/31 11:34:37 Server stopped
```

对文中代码做了改造，代码如下：

```go
var listenAddr string

func init() {
	//接收端口号，默认端口号：5000
	flag.StringVar(&listenAddr, "listen-addr", ":5000", "server listen address")
}
func main() {
	flag.Parse() //外部参数解析
	listenAddr = fmt.Sprintf(":%s",listenAddr)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	//创建 server：
	server := newWebServer(logger)

	done := make(chan struct{}, 1)
	quit := make(chan  os.Signal, 1)

	//os.Interrupt: syscall.SIGINT
	signal.Notify(quit, os.Interrupt)
	//启动另一个 goroutine ，监听将要关闭信号：
	go shutdown(server, logger, quit, done)

	//启动 server：
	logger.Println("Server is ready to handle requests at",listenAddr)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v \n", listenAddr, err)
	}

	//等待已经关闭的信号：
	<-done
	logger.Println("Server stopped")
}

//初始化 server
func newWebServer(logger *log.Logger) *http.Server {
	//路由:
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	//http 服务配置:
	server := &http.Server{
		Addr: listenAddr,
		Handler: router,
		ErrorLog: logger,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	return server
}

//关闭 server
//quit: 接收关闭信号
//done: 发出已经关闭信号
func shutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- struct{}) {
	//等待接收到退出信号：
	<-quit
	logger.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	err := server.Shutdown(ctx)
	if err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v \n", err)
	}

	//do Something ：
	fmt.Println("do something start ..... ", time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println("do something end ..... ", time.Now())

	close(done)
}
```

## 改造的地方

原文代码：

```go
done := make(chan bool, 1)
```

改造后代码：

```go
done := make(chan struct{}, 1)
```

可以看到，把关闭信号从 `chan bool` 改成了 `chan struct{}`。

好处：空 `struct{}` 不占用空间，`chan bool` 占1字节。空 `struct{}` 更节省资源。

由下面测试代码可以知道：

```go
s := struct {}{}
b := true
fmt.Println("struct: ", unsafe.Sizeof(s), " bool: ", unsafe.Sizeof(b)) 
//打印结果：struct:  0  bool:  1
```

## 一些其他
### Go中的信号

`os/signal` 包对信号处理。

- 监听收到的信号：

```go
func Notify(c chan<- os.Signal, sig ...os.Signal)
```

- 取消监听：

```go
func Stop(c chan<- os.Signal)
```

- 监听 Interrupt 信号，用户 `Ctrl+C` 触发：

```go
quit := make(chan  os.Signal, 1)
//os.Interrupt: syscall.SIGINT
signal.Notify(quit, os.Interrupt)	
```

- 监听所有信号：

```go
c := make(chan os.Signal)
signal.Notify(c)
```

- 监听多个指定信号：

```go
c := make(chan os.Signal)
signal.Notify(c,os.Interrupt, os.Kill, syscall.SIGQUIT)
```

- 一些信号：

```go
//操作系统收到信号后的动作:
//Term: 表明默认动作为终止进程
//Ign: 表明默认动作为忽略该信号
//Core: 表明默认动作为终止进程同时输出core dump
//Stop: 表明默认动作为停止进程

// Signals
const (
	SIGABRT   = Signal(0x6) //调用abort函数触发，十进制值：6, Core
	SIGALRM   = Signal(0xe) //时钟定时信号，十进制值：14, Term
	SIGBUS    = Signal(0xa) //非法地址(内存地址对齐错误)，十进制值：10 Core
	SIGCHLD   = Signal(0x14)//子进程结束(由父进程接收)，十进制值：20  Ign
	SIGCONT   = Signal(0x13)//继续执行已经停止的进程(不能被阻塞)，十进制：19 Cont
	SIGEMT    = Signal(0x7)
	SIGFPE    = Signal(0x8)//算术运行错误(浮点运算错误、除数为零等)，十进制：8  Core
	SIGHUP    = Signal(0x1)//终端控制进程结束(终端连接断开),十进制：1  Term
	SIGILL    = Signal(0x4)//非法指令(程序错误、试图执行数据段、栈溢出等)  Core
	SIGINFO   = Signal(0x1d)
	SIGINT    = Signal(0x2)//用户发送INTR字符(Ctrl+C)触发，十进制值：2
	SIGIO     = Signal(0x17)
	SIGIOT    = Signal(0x6)
	SIGKILL   = Signal(0x9)//无条件结束程序(不能被捕获、阻塞或忽略)十进制：9
	SIGPIPE   = Signal(0xd)//消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
	SIGPROF   = Signal(0x1b)
	SIGQUIT   = Signal(0x3)//用户发送QUIT字符(Ctrl+/)触发，十进制值：3
	SIGSEGV   = Signal(0xb)//无效内存引用(试图访问不属于自己的内存空间、对只读内存空间进行写操作)
	SIGSTOP   = Signal(0x11)//停止进程(不能被捕获、阻塞或忽略)
	SIGSYS    = Signal(0xc)
	SIGTERM   = Signal(0xf)
	SIGTRAP   = Signal(0x5)
	SIGTSTP   = Signal(0x12)//停止进程(可以被捕获、阻塞或忽略)
	SIGTTIN   = Signal(0x15)//后台程序从终端中读取数据时触发
	SIGTTOU   = Signal(0x16)//后台程序向终端中写数据时触发
	SIGURG    = Signal(0x10)
	SIGUSR1   = Signal(0x1e)
	SIGUSR2   = Signal(0x1f)
	SIGVTALRM = Signal(0x1a)
	SIGWINCH  = Signal(0x1c)
	SIGXCPU   = Signal(0x18)//超过CPU时间资源限制(4.2BSD)
	SIGXFSZ   = Signal(0x19)//超过文件大小资源限制(4.2BSD)
)
```




