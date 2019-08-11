
## mysql 获取配置信息路径

- 命令行

```
mysqld_safe --datadir=/data/sql_data
```

- 配置文件

```
mysqld --help --verbose | grep -A 1 'Default options'
```

```
Default options are read from the following files in the given order:
/etc/my.cnf /etc/mysql/my.cnf /usr/etc/my.cnf /usr/local/mysql/my.cnf ~/.my.cnf
```

## mysql 配置参数的作用域

- 全局参数

```
mysql> set global 参数名=参数值;
mysql> set @@global.参数名:=参数值;
```

全局参数配置完，需要重新退出才会生效。


- 会话参数

```
mysql> set [session] 参数名=参数值;
mysql> set @@session.参数名:=参数值;
```


## 内存配置相关的参数

- 确定可以使用的内存上限
    - 32位操作系统只能使用3g以内的内存
    - 设置的不能超过物理内存大小

- 确定mysql的每个连接使用的内存

`sort_buffer_size` 

`join_buffer_size`

`read_buffer_size`

`read_rnd_buffer_size`

- 确定需要为操作系统保留多少内存

- 如何为缓存池分配内存

Innodb_buffer_pool_size

总内存 -（每个线程所需要的内存 * 连接数） - 系统保留内存

key_buffer_size

select sum(index_length) from information_schema.tables where engine='myisam'


## IO配置相关的参数





## 安全相关配置参数



##


# 数据库结构优化
良好的数据库逻辑设计 和 物理设计
使查询语句尽量简单

- 尽量减少数据冗余
- 尽量避免数据维护中出现更新，插入和删除异常
    - 插入异常：如果表中的某个实体随着另一个实体而存在。
    - 更新异常：如果更改表中的某个实体的单独属性时，需要对多行进行更新。
    - 删除异常：如果删除表中的某一实体则会导致其他实体的消失
- 节约数据存储空间
- 提高查询效率


- 设计步骤：
    - 需求分析：全面了解产品设计的存储需求
        - 存储需求
        - 数据处理需求
        - 数据的安全性和完整性
    - 逻辑设计：设计数据的逻辑存储结构
        - 数据实体之间的逻辑关系，解决数据冗余和数据维护异常
