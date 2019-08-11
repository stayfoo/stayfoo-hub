## mysql 体系架构

![image](./asset/mysql体系架构图.jpg)

插件式存储引擎

## 存储引擎：MyISAM

MySQL5.5之前版本默认存储引擎

会把索引缓存到内存中，数据缓存到系统中。

- 使用 MyISAM 存储引擎的表： 
    - 系统表
    - 临时表：在排序、分组等操作中，当数量超过一定的大小之后，由查询优化器所建立的临时表。

### 创建一个 MyISAM 表

MyISAM 存储引擎表由 MYD 和 MYI 组成。创建一张 MyISAM 表：

```sql
create table good (id int(11) default NULL, name varchar(10) default NULL) engine=MyISAM default charset=utf8;
```

```bash
#进入到mysql 数据存储目录，/usr/local/mysql/data，对应的数据库目录下： 执行 ls -1
ls -1
good.MYD 
good.MYI
good_5008.sdi
```

- `.MYD` : 存储数据信息
- `.MYI` : 存储索引信息

### MyISAM 特性

- 并发性与锁级别
    - 表级锁（不是行级锁）
        - 对表中的数据修改时，需要对整个表加锁；对表中的数据进行读取的时候，也需要对整个表进行加共享锁（共享锁之间不会阻塞）。（读写和写入是互斥的，在一些情况在对表中数据进行读取的时候，可以在表末尾插入数据。对读写并发不是很好。）

- 表损坏修复（可能会导致数据丢失，不支持事务）

```bash
#对表进行检查
check table tablename
#对表恢复
repair table tablename
```

或者使用命令行 `myisamchk` 对表进行修复，需要mysql停止服务，否则表可能造成更大损坏。

- MyISAM 表支持的索引类型（支持全文索引）
- MyISAM 表支持数据压缩 （压缩后的表只能进行读操作，不能写）
命令行： `myisampack` 


- 限制：
    - 版本 < MySQL5.0时，默认表大小为4G；如果存储大表则要修改 `MAX_Rows` 和 `AVG_ROW_LENGTH`
    - 版本 > MySQL5.0时，默认支持为256TB；

- 适用场景：
    - 非事务型应用（在线分析、报表，不涉及财务的应用）
    - 只读类应用（只读报表等）
    - 空间类应用（存储GPS类数据，支持空间函数）


## 存储引擎：InnoDB

MySQL5.5及以后默认存储引擎。
MySQL5.7以后 Innodb 支持全文索引和空间函数。
Innodb 适合大多数 OLTP 应用。

InnoDB 是事务级存储引擎，完美支持行级锁，事务ACID特性。
会同时在内存中缓存索引和数据。

具有在线热备份方案。

InnoDB 使用表空间进行数据存储。

```
innodb_file_per_table
ON : 数据会存储在独立表空间（ tablename.ibd ）
OFF : 数据会存储在系统表空间（ ibdataX ）
```

```
#查看
show variables like 'innodb_file_per_table'
#设置
set global innodb_file_per_table=off
```


- 系统表空间和独立表空间要如何选择：
    - 系统表空间无法简单的收缩文件大小（删除数据不会缩小；只能把整个数据库innodb表导出后，删除innodb相关的表空间文件后，重启mysql，表空间重建，重新导入数据）
    - 独立表空间可以通过 `optimize table` 命令收缩系统文件
    - 系统表空间会产生IO瓶颈
    - 独立表空间可以同时向多个文件刷新数据
    - 建议：对InnoDB使用独立表空间（mysql5.6之后默认是独立表空间）

- 把系统表空间中的表转移到独立表空间步骤：
    - 1、使用 mysqldump 导出所有数据库表数据（如果数据存储过程有事件触发器一起导出）
    - 2、停止mysql服务，修改参数，并删除innodb相关文件
    - 3、重启mysql服务，重建innodb系统表空间
    - 4、重新导入数据


- Innodb状态检查

```
mysql> show engine innodb status
```
