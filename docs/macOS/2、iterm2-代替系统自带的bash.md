# iterm2

终端神器

## 安装

```bash
brew cask install iterm2
```

- 安装声明高亮

```bash
brew install zsh-syntax-highlighting
```

安装成功之后，在 `~/.zshrc` 最后一行中增加配置：

```bash
vim ~/.zshrc
```

```bash
source /usr/local/share/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
```

- 官网文档：
    - https://www.iterm2.com/documentation-smart-selection.html
- 功能展示：
    - https://www.iterm2.com/features.html

- 主题：
    https://ethanschoonover.com/solarized


## 快捷键

-  Preferences 设置：
    - 按 `Command + ,` 键  ：打开 `Preferences` 配置界面

- 标签操作：
    - `command + t`  新建标签
    - `command + w`   关闭标签
    - `command + 数字 command + 左右方向键`  切换标签
    - `command + enter`	 切换全屏
    - `command + 方向键 → or ←`  切换屏幕

- 分屏：
    - `command + d` 	垂直分屏
    - `command + shift + d`	  水平分屏

- 编辑
    - `command + f` 	查找，支持正则查找
    - `command + ;` 	查看历史命令
    - `command + shift + h` 	查看剪贴板历史
    - `ctrl + u`  清除当前行
    - `ctrl + l`  清屏
    - `ctrl + a`  到行首
    - `ctrl + e`  到行尾
    - `ctrl + f/b`	前进后退

- 历史命令
    - `ctrl + p`	上一条命令
    - `ctrl + r`	搜索命令历史

- 鼠标/触控板操作：
    - iTerm2 中选择文本有三种方式，分别是：
        - 双击：选中单词
        - 三击：选中整行
        - 四击：智能选择
    - iTerm2 默认选中即复制
    - 按下 `Command + Option` 还可以选中矩形范围内的文本

- 快速编辑: `^old^new` 它表示把原来命令(上一条命令)中匹配 `old` 的部分替换成 `new`

```bash
$:(master) ✗ ca dd
zsh: command not found: ca
$: ✗ ^ca^cat
$: ✗ cat dd
cat: dd: No such file or directory
```

- 快速进入 `vim` 编辑:  
    - `Ctrl + x` 然后 `Ctrl + e` 


## iTerm2 配置代理

编辑 `~/.zshrc` : 
```
vim ~/.zshrc
```

增加下面配置（使用的 `shadowsocks`）：

```
proxy list
alias proxy='export all_proxy=socks5://127.0.0.1:1086'
alias unproxy='unset all_proxy'
iTerm2 需要新建标签页，才有效果：

$ proxy
$ curl ip.cn
当前 IP：185.225.14.5 来自：美国

$ unproxy
$ curl ip.cn
当前 IP：115.236.186.130 来自：浙江省杭州市 电信
我们可以测试下：

$ curl https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
```

