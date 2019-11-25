## mysqldump 

- 可以对 msyql 数据库备份
- 可以导出一个数据库表结构
- 可以导出数据库存储的数据
- 可以把导出的数据库表结构和数据库数据，导入到另一个数据库（或另一台机器的另一个数据库）


### 导出&导入数据库表结构

```
# 导出数据库表结构（导出的是创建表结构的sql语句）
mysqldump --opt -d app -uroot -p123 > /data/sql_data/app_struct.sql

# mysqldump --opt -d [数据库名字] -u[数据库用户名] -p[数据库密码] > [导出的文件名字存储路径]
```

- 导入之前需要创建数据库。

```
# 导入数据库表结构
mysql> source /Users/user/Desktop/app_struct.sql;

# mysql> source [要导入的文件名字存储路径];
```

### 导出&导入数据库的数据

```
#导出数据不导出结构
mysqldump -t [数据库名] -u[数据库用户名] -p[数据库密码] > xxx.sql
```

- 导入之前需要创建数据库。

```
# 导入数据库的数据
# mysql> source [要导入的文件名字存储路径];
```



