# Ngnix 环境配置

## 源码安装 Nginx

#### 1、下载Nginx源码包

```
wget http://nginx.org/download/......
```

#### 2、解压：

```
tar -xzf nginx....tar.gz
```

#### 3、配置安装参数 （要添加编译的模块）

- 配置安装路径：

> 配置安装路径，可以指定任意路径，不指定默认安装路径(`/usr/local/nginx`)，一般会自己指定安装目录.

```
--prefix=/home/stayfoo/nginx
```

- 指定编译要增加的模块

```
--with-xxxxxx
```

- 配置命令会生成中间文件，存放在目录为: `/objs` 

>配置完参数，在文件 `/objs/.ngx_modules.c` 中可以查看编译配置包含了哪些 `module`。

- 配置安装参数命令：

```bash
./configure --prefix=/home/stayfoo/nginx \
--with-http_auth_request_module \
--with-http_realip_module \
--with-http_v2_module \
--with-debug \
--with-http_random_index_module \
--with-http_sub_module \
--with-http_addition_module \
--with-http_secure_link_module \
--with-http_geoip_module \
--with-http_ssl_module \
--with-stream_ssl_module \
--with-stream_realip_module \
--with-stream_ssl_preread_module \
--with-stream \
--with-http_slice_module \
--with-google_perftools_module \
--with-threads \
--with-http_gzip_static_module \
--with-http_gunzip_module \
--add-module=/home/web/nginx-http-concat/ \
--add-module=/home/web/ngx_cache_purge/ \
```

上面这些模块有的需要依赖库，所以需要先安装库文件，否则会报错：

```bash
yum -y install gcc gcc-c++ autoconf automake make 

yum install gcc gcc-c++ autoconf automake  zlib zlib-devel openssl openssl-devel pcre-devel

yum -y install pcre-devel

yum install -y zlib-devel

yum -y install openssl openssl-devel

yum -y install GeoIP GeoIP-devel GeoIP-data
```


- 遇到的依赖库报错：

```
//1.
checking for OS + Linux 3.10.0-514.26.2.el7.x86_64 x86_64checking for C compiler ... not found./configure: error: C compiler cc is not found

安装编译工具：
yum -y install gcc gcc-c++ autoconf automake make 
yum install gcc gcc-c++ autoconf automake  zlib zlib-devel openssl openssl-devel pcre-devel
```

```
//2.
./configure: error: the HTTP rewrite module requires the PCRE library.You can either disable the module by using --without-http_rewrite_moduleoption, or install the PCRE library into the system, or build the PCRE librarystatically from the source with nginx by using --with-pcre=<path> option.

yum -y install pcre-devel
```

```
//3.
./configure: error: the HTTP gzip module requires the zlib library.You can either disable the module by using --without-http_gzip_moduleoption, or install the zlib library into the system, or build the zlib librarystatically from the source with nginx by using --with-zlib=<path> option.

yum install -y zlib-devel
```

```
//4.
配置安装目录有问题：(命令  \  换行要有空格，不然命令会连在一起)
Configuration summary  + using system PCRE library  + OpenSSL library is not used  + using system zlib library  nginx path prefix: "/home/stayfoo/nginx--with-http_auth_request_module--with-http_realip_module--with-http_v2_module--with-debug--with-http_random_index_module--with-http_sub_module--with-http_addition_module--with-http_secure_link_module--with-http_geoip_module--with-http_ssl_module--with-stream_ssl_module--with-stream_realip_module--with-stream_ssl_preread_module--with-stream--with-http_slice_module--with-google_perftools_module--with-threads--with-http_gzip_static_module--with-http_gunzip_module--add-module=/home/web/nginx-http-concat/--add-module=/home/web/ngx_cache_purge/"  nginx http client request body temporary files: "client_body_temp"  nginx http proxy temporary files: "proxy_temp"  nginx http fastcgi temporary files: "fastcgi_temp"  nginx http uwsgi temporary files: "uwsgi_temp"  nginx http scgi temporary files: "scgi_temp"
```

```
//5.
./configure: error: SSL modules require the OpenSSL library.You can either do not enable the modules, or install the OpenSSL libraryinto the system, or build the OpenSSL library statically from the sourcewith nginx by using --with-openssl=<path> option.

yum -y install openssl openssl-devel
```

```
//6.
./configure: error: the GeoIP module requires the GeoIP library.You can either do not enable the module or install the library.

yum -y install GeoIP GeoIP-devel GeoIP-data
```

```
//7.
./configure: error: the Google perftools module requires the Google perftoolslibrary. You can either do not enable the module or install the library.

yum install gperftools
```

```
//8.
checking for OS + Linux 3.10.0-514.26.2.el7.x86_64 x86_64checking for C compiler ... found + using GNU C compiler + gcc version: 4.8.5 20150623 (Red Hat 4.8.5-36) (GCC)checking for gcc -pipe switch ... foundchecking for --with-ld-opt="-ltcmalloc" ... not found./configure: error: the invalid value in --with-ld-opt="-ltcmalloc"

```


#### 4、编译Nginx

配置完参数，执行编译

```
make
```

> c 编译生成的所有中间文件放在src目录下

#### 5、安装

```
make install
```

> 首次安装可以使用这个命令。

> 如果是升级安装，需要把编译后的目标文件（`/objs/src`）copy 到安装目录中。


#### 6、启动

- 启动 nginx服务

```
nginx/sbin/nginx
```

- 停止服务

```
nginx/sbin/nginx -s stop
```

- 查看启动情况

```
ps -ef | grep nginx
```


```bash
#查看是否启动成功
curl 192.168.100.10

#查看端口情况
netstat -ano | grep 80
```


#### 7、重启服务

- 检查配置文件是否正确

```
# -t 检查配置文件是否正确
# 默认配置文件 conf/nginx.conf，-c 指定
./sbin/nginx -t -c conf/nginx.conf

nginx: the configuration file /usr/local/nginx-1.12.1/conf/nginx.conf syntax is ok
nginx: configuration file /usr/local/nginx-1.12.1/conf/nginx.conf test is successful
```

- 重启 ngnix

```
./sbin/nginx -s reload
# 重新启动，不会改变启动时指定的配置文件
./sbin/nginx -s reload -c conf/nginx.conf
```

`Nginx` 配置 `vim` 编辑器语法高亮：
拷贝 `Nginx` 源码目录下 `contrib/vim` 目录的所有文件到 `~/.vim/` 目录

```
cp -r contrib/vim ~/.vim
```
