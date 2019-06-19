# mac 设置环境变量

Mac环境变量分为系统级别的配置和用户配置。

- `/etc/`：是系统级别的配置。

- `~/`：是用户级别的配置。

- 配置文件不存在的可以自己新建，加载顺序：

```
1. /etc/profile 
2. /etc/paths
3. ~/.bash_profile
4. ~/.bash_login
5. ~/.profile
6. ~/.bashrc
# 标号即为加载顺序。
```

- 配置文件修改完成需要执行 source：

```
source ~/.bashrc
```

> 一般在 `~/.bashrc` 中配置环境变量
