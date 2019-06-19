# SSH 的配置与管理

## 密码登录

> `SSH` 是 `Secure Shell` 的缩写，其实就是远程 Shell 登录。

- 登录需要指定地址和用户名:

```bash
#登录 IP 地址为 1.1.1.1 的主机上的 root 用户：
ssh root@1.1.1.1
# 回车后输入密码，回车
```

- 默认的 SSH 协议端口号是 `22`, 不需填写

```bash
#有些设备为了安全，会更改 SSH 的端口号，比如改成 1111 ：
ssh -p 1111 root@1.1.1.1
# 回车后输入密码，回车
```

## 免密登陆

`SSH` 支持利用 `RSA` 的公钥和私钥体系来验证身份。

一般在自己电脑生成`RSA`公私钥，然后配置公钥到服务器，也可以了；或者使用工具：`ssh-copy-id`

- 安装工具 `ssh-copy-id`

```bash
brew install ssh-copy-id
```

- 把自己电脑上的公钥复制到目标服务器：

> 自己电脑上的公钥：默认使用 `~/.ssh/id_rsa.pub` 这个文件中的内容
>
> 复制到目标服务器：`~/.ssh/authorized_keys` 文件内

```bash
ssh-copy-id root@1.2.3.4

# 等价于下面这个命令，省掉了两个默认参数
# ssh-copy-id -i ~/.ssh/id_rsa.pub root@100.100.100.100 -p 22
```

## 配置别名

在 `SSH` 的配置文件中给你要登录的服务器（用户名、IP地址、端口号）配置别名。

- 编辑 `~/.ssh/config` 文件：

eg: 给地址为 1.1.1.1，端口为 22，用户名为 root 的远程主机起了一个别名 serverhost

```bash
Host serverhost
    HostName 1.1.1.1
    User root
    Port 22
    IdentityFile ~/.ssh/id_rsa
```

连接服务器：

```
ssh serverhost
```

## 全局 SSH 配置

给 `ssh` 添加一些全局配置，让它变得更好用：

eg：

```
Host *
    ForwardAgent yes
    ServerAliveInterval 10
    ServerAliveCountMax 10000
    TCPKeepAlive no
    ControlMaster auto
    ControlPath ~/.ssh/%h-%p-%r
    ControlPersist 4h
    Compression yes
```

- `ForwardAgent`：
    - 通过 `ssh` 连接到跳板机（服务器A），然后在跳板机（服务器A）链接到服务器B（由跳板机管理的机器），如果不设置 `ForwardAgent` 为 `YES`， 可能会报公钥无效；设置成 `YES` 之后，在两台服务器之间传输数据就不会经过本机。


- `ServerAliveInterval` 和 `ServerAliveCountMax`：
    - 客户端定期向服务端发送心跳包，使得服务端不会断开 `ssh` 的连接。这里表示 10 秒发一次，发 1W 次。

- `ControlMaster` 、`ControlPath`、`ControlPersist`：
    - 每次建立远程连接，在 `~/.ssh` 目录下都会建立一个 `socket` 文件。上面配置表示缓存 `socket` 文件，并保留 4 小时，这样可以加快下次连接的速度。

