# rz & sz 

## centOS平台

- 安装 rz & sz

```
yum install lrzsz
```

- 从服务端发送文件到客户端

```
#sz 要发送的文件
sz filename
```

- 从客户端上传文件到服务端

```
# rz 之后选择要上传的文件
rz
```

## mac 平台 

- 安装lrzsz，使用brew命令：

```
brew install lrzsz
```

如果找不到lrzsz，使用以下命令更新brew库：

```
brew update
```

- 下载 `zmoden` 脚本

    - 下载地址：`https://github.com/mmastrac/iterm2-zmodem`

> 在 `https://github.com/mmastrac/iterm2-zmodem` 上将 `iterm2-send-zmodem.sh` 和 `iterm2-recv-zmodem.sh` 脚本下载下来并放到 `/usr/local/bin/` 目录下，注意赋予脚本执行的权限。


- 配置 iterm2 Trigger 

> 配置 `iterm2` , 让 `iterm2` 支持 `rz/sz` 命令。
>
>1、 打开 `iterm2`
>
>2、 打开 preferences
>
>3、 打开 Profiles -> Default -> Advanced -> Triggers 的Edit按钮
>
>4、 添加配置参数如下：


|Regular Expression | Action | Action|
|-------------------|--------|-------|
| `\*\*B0100` |Run Silent Coprocess | /usr/local/bin/iterm2-send-zmodem.sh |
| `\*\*B00000000000000` |Run Silent Coprocess |/usr/local/bin/iterm2-recv-zmodem.sh|


```
    Regular expression: \*\*B0100
    Action: Run Silent Coprocess
    Parameters: /usr/local/bin/iterm2-send-zmodem.sh

    Regular expression:\*\*B00000000000000
    Action: Run Silent Coprocess
    Parameters: /usr/local/bin/iterm2-recv-zmodem.sh
```

![image](../../asset/zsh-rz-sz.png)

