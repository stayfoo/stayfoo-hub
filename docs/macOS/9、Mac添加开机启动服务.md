# Mac添加开机启动服务

## 偏好设置配置

- 系统偏好设置 -> 用户与群组 -> 登录项

https://blog.csdn.net/guojin08/article/details/19925321

## 配置启动脚本

`~/Library/LaunchAgents`

`LaunchOnlyOnce`： 表示脚本只会执行一次，否则可能会每隔 `10s` 就执行一次。

`Label`：是自定义的名字，理论上来说随便写，不过我建议和文件名保持一致即可，一定不会出错。

`key` 值 `ProgramArguments` 对应
`array`：标签里面就是自己要执行的命令了，前两行 `zsh -c` 不要动，表示用 `zsh`  来执行脚本，然后后面写脚本的路径。比如在这个例子里面我就把入口收敛到了自己的 `onlogin.sh` 里面，然后再执行任何事情就很方便了。



```bash
# 将这个文件注册到系统中。
sudo launchctl load ~/Library/LaunchAgents/com.mengyueping.onlogin.plist

# 报错，权限不够：
/Users/mengyueping/Library/LaunchAgents/com.mengyueping.onlogin.plist: Path had bad ownership/permissions

# 添加权限
sudo chmod 644 ~/Library/LaunchAgents/com.mengyueping.onlogin.plist 
# 查看权限，同级目录下的其他文件权限
ls -l /Users/mengyueping/Library/LaunchAgents/com.mengyueping.onlogin.plist
```
