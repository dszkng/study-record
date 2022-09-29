[toc]

# 基本命令

```
docker images

docker search codingcorp-php:1.0

docker run -it -p 80:80 -v ./project:/var/www/html --name=codingcorp-php1 codingcorp-php:1.0
docker run -itd -p 80:80 -v ./project:/var/www/html --name=codingcorp-php1 codingcorp-php:1.0
docker run -it -p 80:80 -v ./project:/var/www/html --name=codingcorp-php1 codingcorp-php:1.0 php bin/laravels start

docker exec -it codingcorp-php1 bash

#创建一个新的容器但不启动它
docker create --name codingcorp-php1 codingcorp-php:1.0

docker stop codingcorp-php1

docker start codingcorp-php1

docker restart codingcorp-php1

docker commit codingcorp-php1 codingcorp-php:1.1
docker commit -a "dszkng" -m "save as new version" codingcorp-php1 codingcorp-php:1.1

docker cp xx.png codingcorp-php1:/var/www/html/php/public/images/
docker cp codingcorp-php1:/var/www/html/php/public/images/xx.png ./

docker rm codingcorp-php1

docker rmi codingcorp-php:1.0

#暂停容器中所有进程
docker pause codingcorp-php1
docker unpause codingcorp-php1

#杀掉容器的进程
docker kill codingcorp-php1

#向容器内部的主进程发送 SIGKILL 信号（默认），或使用 --signal 选项指定的信号号
docker kill -s KILL codingcorp-php1

#获取容器的日志：跟踪日志输出
docker logs -f codingcorp-php1

#仅列出最新N条容器日志
docker logs --tail=10 codingcorp-php1

#显示某个开始时间的所有日志
docker logs --since="2022-05-01" codingcorp-php1
docker logs --since="2022-05-01" --tail=10 codingcorp-php1
```

```
#列出容器：已运行、所有、最近创建、显示文件大小、查询指定条数
docker ps
docker ps -a
docker ps -l
docker ps -s
docker ps -n 5
```

```
#7种状态
created（已创建）
restarting（重启中）
running（运行中）
removing（迁移中）
paused（暂停）
exited（停止）
dead（死亡）
```

```
#查看容器中运行的进程信息，支持 ps 命令参数
docker top codingcorp-php1
docker top codingcorp-php1 | grep laravels

#容器运行时不一定有/bin/bash终端来交互执行top命令，而且容器还不一定有top命令，可以使用docker top来实现查看container中正在运行的进程。
```

```
#获取容器/镜像的元数据
docker inspect codingcorp-php:1.0
```

```
#连接到正在运行中的容器
docker attach codingcorp-php1
#确保CTRL-D或CTRL-C不会关闭容器
docker attach --sig-proxy=false codingcorp-php1
```

```
#从服务器获取实时事件
docker events --since="1467302400"
docker events -f "image"="codingcorp-php:1.0" --since="2022-05-01"
```

```
#阻塞运行直到容器停止，然后打印出它的退出代码
docker wait codingcorp-php1
```

```
#列出指定的容器的端口映射，或者查找将PRIVATE_PORT NAT到面向公众的端口
docker port codingcorp-php1
```

```
#查看指定镜像的创建历史
docker history codingcorp-php:1.0
#-H表示以可读的格式打印镜像大小和日期，默认为true
docker history -H codingcorp-php:1.0
```

```
#检查容器里文件结构的更改
docker diff codingcorp-php1
```

```
#显示 Docker 系统信息，包括镜像和容器数
docker info

#显示 Docker 版本信息
docker version
```

# 导入导出

## load和save

1. docker load:将镜像存储文件导入到本地镜像库；
2. docker save:将一个镜像导出为文件，保存的是该镜像的所有历史记录；

```
docker load < codingcorp-php.tar.gz
docker load --input codingcorp-php.tar

docker save -o codingcorp-php.tar uustudy/codingcorp-php:1.0
```

## import和export

1. docker import:导入一个容器快照到本地镜像库；

2. docker export:将一个容器导出为文件，保存的是容器当时的状态，即容器快照；

```
docker export 1e560fca3906 > codingcorp-php.tar

docker import codingcorp-php.tar uustudy/codingcorp-php:1.0
docker import http://example.com/codingcorp-php.tgz uustudy/codingcorp-php:1.0
```

## 区别和联系

docker save和docker export之间的区别：

1. docker save是将镜像保存为tar包，且会保存该镜像的父层、标签、所有历史等信息；docker export是将容器文件系统保存为tar包，保存的是容器当时的状态（快照）；
2. docker save可以同时指定多个镜像名称；docker export只能指定一个容器名称；
3. docker save保存的镜像文件tar包使用docker load命令加载还原；docker export保存的容器快照tar包使用docker import命令导入还原；
4. docker save保存的tar包文件通常比docker export导出的文件要大；

docker load和docker import之间的区别：

1. docker load将镜像存储文件导入到本地镜像库；docker import将容器快照文件导入到本地镜像库；
2. docker load不能指定url；而docker import可以指定url来进行导入；

save=打包一组镜像
load=导入一组镜像
export=将容器导出成一个镜像
import=将镜像加载成一个容器

# Dockerfile

Dockerfile 是一个用来构建镜像的文本文件，文本内容包含了一条条构建镜像所需的指令和说明。

shell 模式和 exec 模式：

```
#shell
RUN <命令行命令>
RUN /usr/local/bin/php think run -p 8080

#exec
RUN ["可执行文件", "参数1", "参数2"]
RUN ["/usr/local/bin/php", "think", "run", "-p", "8080"]
RUN ["/usr/local/bin/php", "think", "run", "-p=8080"]
```

## 语法格式

```
MAINTAINER

#指令用来给镜像添加一些元数据（metadata），以键值对的形式
LABEL <key>=<value> <key>=<value> <key>=<value> ...
LABEL org.opencontainers.image.authors="runoob"

FROM

EXEC

#每一个 RUN 命令都是新建的一层
RUN

#运行程序，和RUN二者运行的时间点不同：CMD 在docker run 时运行；RUN 是在 docker build。
#CMD 指令指定的程序可被 docker run 命令行参数中指定要运行的程序所覆盖。
#注意：如果 Dockerfile 中如果存在多个 CMD 指令，仅最后一个生效。
CMD
CMD ["<param1>","<param2>",...]  # 该写法是为 ENTRYPOINT 指令指定的程序提供默认参数

EXPOSE

WORKDIR

ADD

COPY

#构建参数，与 ENV 作用一致。不过作用域不一样。ARG 设置的环境变量仅对 Dockerfile 内有效，也就是说只有 docker build 的过程中有效，构建好的镜像内不存在此环境变量。
ARG

ENV

#挂载点，定义匿名数据卷。在启动容器时忘记挂载数据卷，会自动挂载到匿名卷。
#在启动容器 docker run 的时候，我们可以通过 -v 参数修改挂载点。
VOLUME ["<路径1>", "<路径2>"...]
VOLUME <路径>

#用于指定执行后续命令的用户和用户组，这边只是切换后续命令执行的用户（用户和用户组必须提前已经存在）
USER <用户名>[:<用户组>]

#类似于 CMD 指令，但其不会被 docker run 的命令行参数指定的指令所覆盖，而且这些命令行参数会被当作参数送给 ENTRYPOINT 指令指定的程序。
#但是, 如果运行 docker run 时使用了 --entrypoint 选项，将覆盖 ENTRYPOINT 指令指定的程序。
#注意：如果 Dockerfile 中如果存在多个 ENTRYPOINT 指令，仅最后一个生效。
ENTRYPOINT
ENTRYPOINT ["<executeable>","<param1>","<param2>",...]

#用于指定某个程序或者指令来监控 docker 容器服务的运行状态
#设置检查容器健康状况的命令
HEALTHCHECK [选项] CMD <命令>
#如果基础镜像有健康检查指令，使用这行可以屏蔽掉其健康检查指令
HEALTHCHECK NONE
#这边 CMD 后面跟随的命令使用，可以参考 CMD 的用法
HEALTHCHECK [选项] CMD <命令>

HEALTHCHECK --interval=5s --timeout=5s \
CMD curl -sS 'http://localhost:9200' || exit 1
```

举例：

```
FROM nginx

ENTRYPOINT ["nginx", "-c"] # 定参
CMD ["/etc/nginx/nginx.conf"] # 变参 
```

## 构建

```
docker build -t codingcorp-php:1.0 .
docker build --build-arg "GITHUB_TOKEN=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" . -t codingcorp-php:1.0
```

# 发布

```
docker login
docker login -u dszkng -p xxx

docker logout

docker pull codingcorp-php:1.0
docker pull -a codingcorp-php:1.0

docker tag codingcorp-php:1.0 harbor.oneitfarm.com/uustudy/codingcorp-php:1.0

docker push harbor.oneitfarm.com/uustudy/codingcorp-php:1.0
```

# Compose

安装：https://www.runoob.com/docker/docker-compose.html

docker-compose

# 镜像转换

## vm iso => docker image

将linux虚拟机做成docker镜像

```
#1.将虚拟机打包
 
tar --numeric-owner --exclude=/proc --exclude=/sys -zcvpf debian-3519.tar /
--numeric-owner:执行所属
--exclude：排除那些文件或者目录 
-zcvf ：打包压缩 p保持文件的绝对路径

#2.导入镜像
将打包好的docker镜像放到装有docker的系统上，执行如下命令，导入镜像：

docker import debian-3519.tar(这是上一步打包好的docker镜像) image名字（自己定义）

#3.运行镜像
docker run -it  image名字（上一步自己定义的） /bin/bash
```

## docker image => vm iso

https://www.qedev.com/cloud/226997.html

# 其他

## 命令

查看完整的COMMAND命令

```
docker ps -a --no-trunc
```

查看容器内部IP

```
docker inspect containerid | grep IPAddress
```

## Q&A

1. 容器运行起来就自动退出了

现象：

```
docker run ...

# Exited(0)
docker ps -a
```

原因：缺失一个前台进程。docker的机制问题，容器运行必须有一个前台进程，如果没有前台进程执行，容器认为空闲，就会自行退出。

解决：运行top，tail、循环或运气其他软件的前台进程等等。

```
docker run ... /bin/sh -c "while true;do echo hello;sleep 10000;done"
```

2. scratch 是什么镜像？

docker最小系统镜像

https://blog.csdn.net/HCY_2315/article/details/119899397

3. 有哪些官方的系统基础镜像？

```
ubuntu
debian
centos
fedora
alpine 
```

- Ubuntu：在选择Ubuntu发行版本时要选择开头是偶数且后缀为04的版本，例如16.04，18.04，20.04，22.04
- Debian：适合于服务器的操作系统，它比Ubuntu要稳定得多，整个系统基础核心非常小，不仅稳定，而且占用硬盘空间小，占用内存小。
- CentOS：生产环境我们常用CentOS6、7、8的最新稳定版
- Alpine：一个面向安全的轻型 Linux 发行版，下载速度加快，镜像安全性提高，主机之间的切换更便捷，占用更少磁盘空间，目前 Docker 官方已开始推荐使用 Alpine 替代之前的 Ubuntu 做为基础镜像环境

https://blog.csdn.net/timonium/article/details/119023016

4. 为什么要传上下文路径`.`？

>由于 docker 的运行模式是 C/S。我们本机是 C，docker 引擎是 S。实际的构建过程是在 docker 引擎下完成的，所以这个时候无法用到我们本机的文件。这就需要把我们本机的指定目录下的文件一起打包提供给 docker 引擎使用。

>如果未说明最后一个参数，那么默认上下文路径就是 Dockerfile 所在的位置。

>注意：上下文路径下不要放无用的文件，因为会一起打包发送给 docker 引擎，如果文件过多会造成过程缓慢。


## 坑

1. 未清理安静，导致磁盘占用较多

write /var/lib/docker/overlay2/ugjn6ltq40xj7b50qkosugtrg/diff/extensions/swoole-4.4.12.tgz: 

```
no space left on device
```

docker system df

磁盘空间 or 内存不够

docker system prune -a -f

Caution: Do not run this on a production environment	

Docker启动出现”No space left on device” 或者 docker日志太多导致磁盘占满问题					
https://www.tracymc.cn/archives/2949					

docker 安装 php 扩展
docker-php-ext-install zip		

docker 安装 zip		
apk add libzip-dev	
docker-php-ext-install zip	

2. docker build失败

原因：内存不够，进程被杀了

解决：

```
docker run ... -m=32m ...
docker build NODE_OPTIONS="--max_old_space_size=4096" #8192
```

3. docker 权限问题 Got permission denied while trying to connect to the Docker daemon socket at...

```
在用户权限下docker 命令需要 sudo 否则出现以下问题

通过将用户添加到docker用户组可以将sudo去掉，命令如下

sudo groupadd docker #添加docker用户组

sudo gpasswd -a $USER docker #将登陆用户加入到docker用户组中

newgrp docker #更新用户组

ubuntu18.04在重启后会生效，如果不是特别着急，可以先重启然后再做docker操作。

原文链接：https://blog.csdn.net/u011337602/article/details/104541261/
```