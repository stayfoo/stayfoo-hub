## CentOS系统参数优化

推荐书籍：《Linux性能优化大师》

内核相关参数( `/etc/sysctl.conf` )

- 查看内核(`kernel`)参数配置

```
sysctl -a
```

#### 网络性能参数

```
net.core.somaxconn=65535
net.core.netdev_max_backlog=65535
net.ipv4.tcp_max_syn_backlog=65535
```

http 请求经过三次握手建立网络连接，处于监听状态的端口，都会有自己的监听队列，参数 `net.core.somaxconn` 就决定了监听队列大小的长度，负载很大的服务器，就需要把这个参数修改大一些。

`net.core.netdev_max_backlog` 在每个网络接口接收数据包的速率 比 内核处理数据包的速率快的时候，允许被发送到队列中的数据包的最大数量。

`net.ipv4.tcp_max_syn_backlog` 还未获得对方连接的请求，可以保存到队列中的最大数目，超过这个数据大小的连接可能就会被抛弃。


```
#查看有多少个请求由监听变成了链接

#列出所有的端口，包括监听的和未监听的。
netstat -a

Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State
tcp        0      0 0.0.0.0:etlservicemgr   0.0.0.0:*               LISTEN
```

```
#列出所有的tcp协议的端口
netstat -t

#列出所有的UDP协议的端口
netstat -ua

#找出程序运行的端口
netstat -ap | grep '程序名'

#找出端口的程序名
netstat -ap | grep '端口号'

#显示路由表的信息
netstat -r

#显示接口信息
netstat -i

#分类统计各个协议的相关信息
netstat -sa
```

- `net.ipv4.tcp_fin_timeout` 用于处理 tcp 连接等待状态的时间


```
#用于加快 tcp 连接的回收，tcp连接占满就会出现无法连接的状态
net.ipv4.tcp_fin_timeout=10
net.ipv4.tcp_tw_reuse=1
net.ipv4.tcp_tw_recycle=1
```

- tcp 连接接收和发送缓冲区大小的默认值和最大值（调整的大一些）

```
net.core.wmem_default = 87380
net.core.wmem_max = 16777216
net.core.rmem_default=87380
net.core.rmem_max=16777216
```

- 减少失效连接所占用的 tcp 资源的数量，加快资源回收的效率（调整小一些）

```
#tcp 发送 keepalive 的时间间隔，用户确认tcp连接是否有效，单位秒
net.ipv4.tcp_keepalive_time=120

#当发送探测点消息没有响应时，重发该消息的时间间隔，单位秒
net.ipv4.tcp_keepalive_intvl=30

#认定 tcp 连接失效之前，最多发送多少个 keepalive 消息
net.ipv4.tcp_keepalive_probes=3
```


#### 内存性能参数


```
kernel.shmmax = 4294967295
```

- 这个参数应该设置的足够大，以便能在一个共享内存段下容纳下整个的 Innodb 缓冲池的大小。（如果太低，就需要创建多个共享内存段，可能导致系统性能下降，原因是当实例启动的时候，多个共享内存段可能会导致当时系统性能轻微的性能下降，其他时候不会有影响）
- 这个值的大小对 64 位 linux 系统，可取的最大值为物理内存值-1byte，建议值为大于物理内存的一半，一般取值大于Innodb缓冲池的大小即可，可以取物理内存-1byte。


```
vm.swappiness = 0
```

- 这个参数当内存不足时，会对性能产生比较明显的影响


Linux 系统内存交换区：
在Linux系统安装时，都会有一个特殊的磁盘分区，称之为系统交换分区。
使用 `free -m` 在系统中查看，可以看到类似下面的内容，其中 `swap` 就是交换分区。

```
[root@VM_0_3_centos ~]# free -m
              total        used        free      shared  buff/cache   available
Mem:           3790        1763         235           0        1791        1732
Swap:             0           0           0
```

当操作系统因为没有足够的内存时，就会将一些虚拟内存写到磁盘的交换区中，这样就会发生内存交换。


在MySQL 服务器上是或否要使用交换分区有一些争议：在MySQL服务所在的Linux系统上完全禁用交换分区。
带来的风险：降低操作系统的性能；容易造成内存溢出，崩溃，或都被操作系统kill掉。

在MySQL服务器上保留交换分区还是很有必要的，但是要控制何时使用交换分区。


```
vm.swappiness = 0
```

- 参数告诉Linux内核除非虚拟内存完全满了，否则不要使用交换分区。



#### 增加资源限制 `/etc/security/limit.conf`


这个文件实际上市 Linux PAM 也就是插入式认证模块的配置文件。
打开文件数的限制。


```
* soft nofile 65535
* hard nofile 65535
```

- 用于控制打开文件的数量，加到 limit.conf 文件的末尾即可。 增加到 65535 以保证可以打开足够多的文件句柄。
- * 表示对所有用户有效
- soft 指的是当前系统生效的设置
- hard 表明系统中所有设定的最大值
- nofile 表示所限制的资源是打开文件的最大数目
- 65536 就是限制的数量
- 这个文件修改需要重启系统才能生效。



#### 磁盘调度策略 

`/sys/block/devname/queue/scheduler`

```
cat /sys/block/vda/queue/scheduler

cat /sys/block/sda/queue/scheduler
noop anticipatory deadline [cfq]
```

- noop（电梯式调度策略） ：NOOP 实现了一个 FIFO 队列，它像电梯的工作方法一样对 I/O 请求进行组织，当有一个新的请求到来时，它将请求合并到最近的请求之后，以此来保证请求同一介质。 NOOP 倾向饿死读而立于写，因此 NOOP 对于闪存设备、RAM及嵌入式系统是最好的选择。

- deadline（截止时间调度策略）： deadline 确保了一个截止时间内服务请求，这个截止时间是可调整的，而默认读期限短于写期限。 这样就防止了写操作因为不能被读取而饿死的现象， deadline 对数据库类应用是最好的选择。
- anticipatory （预料I/O调度策略）： 本质上与deadline 一样，但在最后一次读操作后，要等待6ms，才能继续进行对其他I/O请求进行调度。 它会在每个6ms中插入新的I/O操作，而会将一些小写入流合并成一个大写入流，用写入延时换取最大的写入吞吐量。 AS适合于写入较多的环境，比如文件服务器，AS对数据库环境表现很差。






