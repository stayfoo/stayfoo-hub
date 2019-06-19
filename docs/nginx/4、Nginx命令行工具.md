## Nginx命令行

`nginx [cmd]`

- `-?`、`-h` ：帮助
- `-c [配置文件路径]` ：使用指定的配置文件
- `-g` ：指定配置指令
- `-p` ：指定运行目录
- `-s` ：发送信号
    - `stop` ：立刻停止服务
    - `quit` ：优雅的停止服务
    - `reload` ：重载配置文件
    - `reopen` ：重新开始记录日志文件
- `-t`、`-T` ：测试配置文件是否有语法错误
- `-v` 、`-V` ：打印nginx的版本信息、编译信息等


重载配置文件：`nginx -s reload`

热部署：（更换Nginx版本-二进制文件）

```bash
#查看Nginx进程
ps -ef | grep nginx

#备份二进制文件
cp nginx nginx.old

#拷贝新二进制文件
cp -r  nginx /nginx/sbin/ -f

#杀死进程
kill -USR2 [进程号]

#ps查看进程，新旧work进程都在运行
#关闭所有worker进程
kill -WINCH [进程号]
```

切割日志文件：

```bash
#剪切log
mv access.log bak.log

/sbin/nginx -s reopen

```

## Nginx进程管理：信号

使用信号进行进程间管理

- Master 进程：
    - 监控worker进程
        - CHLD（当子进程终止的时候，会向父进程发送CHLD信号）
    - 管理worker进程：通过接收信号
        - TERM，INT（立刻停止请求）
        - QUIT（优雅的停止请求，向用户的tcp请求报文）
        - HUP（重载配置文件）
        - USR1（重新打开日志文件，做日志文件的切割）（上面四个是可以通过Nginx命令行发送的信号）
        - 
        - `USR2` （热部署使用）（只能通过Linux命令行care发送，需要知道master进程的pid）
        - `WINCH`（热部署使用）（只能通过Linux命令行care发送，需要知道master进程的pid）
- Worker进程，也可以接受信号，一般都是通过master进程管理worker进程，不直接管理
    - TERM，INT
    - QUIT
    - USR1
    - WINCH
- nginx命令行
    - reload：HUP
    - reopen：USR1
    - stop：TERM
    - quit：QUIT

启动Nginx之后，Nginx会把它的pid记录在安装目录的logs文件夹下的一个`nginx.pid`文件

