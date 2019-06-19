# go mod

## `go mod` 命令管理包


- 在当前目录初始化生成 `go.mod` 文件

```
go mod init 
```

- 下载包依赖到本地缓存

本地缓存目录：`$GOPATH/pkg/mod`

```
go mod download
```

- 编辑`go.mod`

格式化`go.mod`文件

```
go mod edit -fmt
```

以`json`的形式查看依赖

```
go list -m -json all
go mod edit -json
```

- 打印模块依赖图

```
go mod graph
```

- 拉取缺少的模块，移除不用的模块

```
go mod tidy
```

- 将依赖复制到`vendor`下

```
go mod vendor
```

- 验证依赖是否正确

```
go mod verify
```

- 需要依赖的原因

```
go mod why
```

## `go.mod` 编写

四个命令: 

- `module` : 指定包的名字（路径）
- `require`: 指定的依赖项模块
- `replace`: 可以替换依赖项模块
- `exclude`: 可以忽略依赖项模块

eg:

```go
module cocktail

go 1.12

require github.com/gin-gonic/gin v1.4.0

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/net => github.com/golang/net v0.0.0-20190514140710-3ec191127204
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 => github.com/golang/net v0.0.0-20190514140710-3ec191127204
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c => github.com/golang/net v0.0.0-20190514140710-3ec191127204

	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => github.com/golang/sys v0.0.0-20190516110030-61b9204099cb
	golang.org/x/sys v0.0.0-20190412213103-97732733099d => github.com/golang/sys v0.0.0-20190516110030-61b9204099cb

	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.2
	golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e => github.com/golang/tools v0.0.0-20190517183331-d88f79806bbd
)
```

使用 `replace` 可以把 `golang.rog/` 替换成 `github/ ` 。

## 一些其他使用

- 执行 `go run server.go` 运行代码会发现 `go mod` 会自动查找依赖自动下载。

- `go module` 安装 `package` 的原則是先拉最新的 `release tag`，若无`tag`则拉最新的`commit`。 

- `go` 会自动生成一个 `go.sum` 文件来记录 `dependency tree`。

- 来检查可以升级的 `package`

```
go list -m -u all
```

- 升级依赖版本
```
go get -u need-upgrade-package
```

或者 
```
go get -u
```


> Modules官方介绍：https://github.com/golang/go/wiki/Modules


## 一些遇到的问题

- go 1.10.3 版本Bug。需要升级版本,删除 `/usr/local/Cellar/go/` 目录下文件，重新安装。

原因参考：https://github.com/golang/go/issues?q=milestone%3AGo1.10.4

```bash
go get github.com/beego/bee

# github.com/beego/bee
/usr/local/Cellar/go/1.10.3/libexec/pkg/tool/darwin_amd64/link: /usr/local/Cellar/go/1.10.3/libexec/pkg/tool/darwin_amd64/link: combining dwarf failed: Unknown load command 0x32 (50)
```

- go mod init 报错

```
go mod 

go: modules disabled inside GOPATH/src by GO111MODULE=auto; see 'go help modules'

# 设置环境变量解决 GO111MODULE 默认值为auto  其他：on 、off
export GO111MODULE=on
```

