# mysql安装配置

## Mac系统安装

官网直接下载客户端安装。

https://dev.mysql.com/downloads/mysql/

解决 mac 上 mysql 不能远程访问的问题，可以参考博客： https://www.jianshu.com/p/f2468b40d1e7

## CentOS 7 
### 安装 MySQL

- 1、yum仓库下载MySQL：

```
sudo yum localinstall https://repo.mysql.com//mysql80-community-release-el7-1.noarch.rpm

```

- 2、yum安装MySQL：

```
sudo yum install mysql-community-server
```


- 3、启动MySQL服务：

```
systemctl start mysqld.service
# 或者 
sudo service mysqld start
```

- 4、检查 MySQL 服务状态：

```
sudo service mysqld status
```

- 5、查看初始密码（如无内容直接跳过）：

```
sudo grep 'temporary password' /var/log/mysqld.log
```

- 6、MySQL客户端登录：

```
mysql -uroot -p
```

- 7、输入密码为第5步查出的，如果没有，直接回车，然后输入命令  

```
flush privileges
```

- 8、修改 root 登录密码：

```
#注意：要切换到mysql数据库，使用use mysql
# 密码需要包含：大小写字母、特殊字符、数字
ALTER USER 'root'@'localhost' IDENTIFIED BY '新密码';
```

如果忘记密码，可以通过 `skip-grant-tables` 配置跳过输入密码登录 `MySQL`，执行 7、8 步进行修改，如果`'root'@'localhost'`变为`'root'@'%'`，那么`alter`语句中的也要修改。

- 9、配置 MySQL 允许外部访问：

1）首先设置阿里云安全组规则入方向，支持MySQL端口3306可访问（协议类型下拉菜单中有MySQL端口）; 2）服务端登录 MySQL
    
```
use mysql;
```

然后执行

```
select user,host from user
```

可查看用户及对应允许访问主机，然后执行 

```
update user set host = '%' where user ='root';
```

允许任何外部可访问；再执行上一步查看命令，可比较结果.

- 10、如此即可连接

可查看MySQL服务端口，如果看到的value为0，则说明没有使用密码登录，需要去修改my.cnf文件；

```
show global variables like 'port';
```

`my.cnf` 文件也可以通过 `port=3306` 来指定 MySQL 服务端口，重启 MySQL 服务即可。

## 彻底卸载 MySQL

- 查看 MySQL 是否安装:

```
# 方式一：
yum list installed mysql*  

Loaded plugins: fastestmirror  
Loading mirror speeds from cached hostfile  
 * base: mirrors.yun-idc.com  
 * extras: mirror.neu.edu.cn  
 * updates: mirrors.yun-idc.com  
Installed Packages  
MySQL-client.x86_64   5.6.27-1.el6    installed  
MySQL-devel.x86_64    5.6.27-1.el6    installed  
MySQL-server.x86_64   5.6.27-1.el6    installed  
```

```
#方式二：
rpm -qa | grep -i mysql 

MySQL-server-5.6.27-1.el6.x86_64  
MySQL-client-5.6.27-1.el6.x86_64  
MySQL-devel-5.6.27-1.el6.x86_64  
```

- 卸载

```
rm -rf /var/lib/mysql  

rm /etc/my.cnf  
```

如果装了 `mysql-devel` ，卸载为：

```
yum remove mysql mysql-devel mysql-server mysql-libs compat-mysql51  
```

注（例如）:
mysql-5.5.39-1.el6.remi.x86_64
mysql-libs-5.5.39-1.el6.remi.x86_64
compat-mysql51-5.1.54-1.el6.remi.x86_64
mysql-server-5.5.39-1.el6.remi.x86_64


```
#卸载2   {继续，1，2选择一种（此处为介绍）：}:


[root@localhost mysql]# rpm -aq | grep -i mysql
MySQL-server-5.6.27-1.el6.x86_64
MySQL-client-5.6.27-1.el6.x86_64
MySQL-devel-5.6.27-1.el6.x86_64

[root@localhost mysql]# rpm -e MySQL-server-5.6.27-1.el6.x86_64 
[root@localhost mysql]# rpm -e MySQL-client-5.6.27-1.el6.x86_64  
[root@localhost mysql]# rpm -e MySQL-devel-5.6.27-1.el6.x86_64  
[root@localhost rc.d]# cd /var/lib/
[root@localhost lib]# rm -rf mysql/

```

注: 删除 MySQL 数据库目录(关键) ，否则 password 不更新（默认安装，如果自定义安装路径和链接路径 ln -s …… 要删除。）

```
rm -rf /var/lib/mysql
```

```
# 卸载3：

[root@localhost usr]# whereis mysql  
mysql: /usr/lib64/mysql  

[root@localhost usr]# rm -rf /usr/lib64/mysql  

```

注：清空相关mysql的所有目录以及文件和其他配置和设置等。如果有，则删除。也必须考虑其他软件不去影响。

```
find / -name mysql

rm -rf /usr/lib/mysql
rm -rf /usr/share/mysql
```

```
# 卸载4：

[root@localhost usr]# rm –rf /usr/my.cnf  

[root@localhost usr]# rm -rf /root/.mysql_sercret   
```

```
# 卸载5（自启服务）：

[root@localhost usr]# chkconfig --list | grep -i mysql  

[root@localhost usr]# chkconfig --del mysqld  

# 此处删除看自己设置:mysql/mysqld  
```



