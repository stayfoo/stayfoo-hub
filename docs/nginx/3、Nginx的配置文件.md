# 配置文件

- 配置文件由`指令`与`指令块`构成

## 指令语法

- `{ }` ：将多条指令组织在一起
- `include` ：语句允许组合多个配置文件（提高可维护性）
- `#` ：添加注释 
- `$` ：使用变量
- 部分指令的参数支持正则表达式

配置参数：

|时间的单位  | 全拼|
|---|---|
|`ms`| `milliseconds` |
|`s` | `seconds`|
|`m` | `minutes`|
|`h` | `hours`|
|`d` | `days`|
|`w` | `weeks`|
|`M` | `months, 30 days`|
|`y` | `years,365 days`|


|空间的单位  | 全拼|
|---|---|
|不加任何后缀| `bytes` |
|`k/K` | `kilobytes`|
|`m/M` | `megabytes`|
|`g/G` | `gigabytes`|

## 指令块

配置指令块：`http` `server` `upstream` `location`

### location

```nginx
http {
    
    #定义日志格式
    #$remote_addr:远端的ip地址
    #$time_local:当时时间
    #$status:状态码
    #
    log_format main '$remote_addr - $remote_user [$time_local] "$request" '
                    '$status $body_bytes_sent "$http_referer" '
                    '"$http_user_agent" "$http_x_forwarded_for"';

    access_log logs/my_api.access.log main;

    # alias： 指定url路径和本地的文件目录一一对应
    # root：会把url中的路径带到文件目录中来
    location / {
        alias dlib/;
        set &limit_rate 1k;#每秒传输1k字节到浏览器中（流量限制，防止大文件下载占用带宽）。 set + 内置变量。
        
        gzip on; #gzip开光
        gzip_min_length 1; #文件大小小于1字节就不再执行压缩
        gzip_comp_level 2; #压缩级别
        gzip_types text/plain application/x-javascript text/css application/xml text/javascript image/jpeg image/gif image/png; #指定哪些文件类型进行压缩
    }
}
```

### http模块

配置指令冲突，以谁为准？


```nginx
main
http {
    upstream {}
    split_clients {}
    map {}
    geo {}
    server {
        if(){}
        location {
            limit_except {}
        }
        location {
            location {
                
            }
        }
    }
    server {
        
    }
}
```

指令的 Context

指令的合并：

值指令：存储配置项的值，可以合并，eg: root 、access_log 、gzip

动作类指令：指定行为, 不可以合并，eg: rewrite 、proxy_pass ，生效阶段：server_rewrite 阶段、rewrite阶段、content阶段


值指令的指令继承规则：向上覆盖。子配置不存在时，直接使用父配置块；子配置存在时，直接覆盖父配置块；

```nginx
server {
    listen 8080;
    root /home/geek/nginx/html;
    access_log logs/geek.access.log main;
    location /test {
        root /home/geek/nginx/test;
        access_log logs/access.test.log main;
    }
    location /dlib {
        alias dlib/;
    }
    
    location / {
        
    }
}
```


- listen 指令


```
listen *.80|*:8000;   //Default

listen unix:/var/run/nginx.sock;
listen 127.0.0.1:8000;
listen 127.0.0.1;
listen 8000;
listen *:8000;
listen localhost:8000 bind;
listen [::]:8000 ipv6only=on;
listen [::1];
```

