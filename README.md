# site-demo

### 这是一个半成品

实现了热编译, go server 和 Mongo, ES链接好了, 但是最底层的API还没有写

## 期望

搭建一个基于docker容器的网站demo, 为了方便, 先尝试使用beego,计划使用MySQL存放常用读写数据, 使用MongoDB存储常读少改的数据,使用Elasticsearch做索引
前端...实践下React

## TODO

- livereload

  目前热编译用的很不舒服, 有时间实现下livereload...或者看看改下bee工具
- 研究下要不要使用`gocrawl`
- 学一下`MongoDB` = =
- 爬取试验数据存入`MongoDB`一遍后续使用`Elasticsearch`索引

# install

## 说明

因为`golang`的编译器在容器内,容器内只能看到`volume`范围的东西, 所以要用到`vedor`把包依赖都整理起来放到`volume`范围内,
这样容器内的程序才能正常编译运行,

我弄了个`gin` 来做热编译的事情, 感觉没有`bee`那么好用, 但是gin的扩展性要比`bee`好, 以后改起来方便

但是日常创建项目还是完全可以使用`bee`的


### container 设置:

- 到 docker 目录下根据Makefile制作image
- 修改docker-compose.yml文件

  注意`working_dir`必须要在`$GOPATH`下

- 习惯设置`docker-compose` alias 为 `fig`:
>alias fig=docker-compose

### 保存`go`代码

- 在根目录使用`bee`工具创建一个新的项目(或者手动创建)
- 项目能正常运行之后, 在根目录执行

>make vendor
保存依赖包

## 注意

使用`vendor`需要安装`vendor`工具,我用的是godep:

```
github.com/tools/godep
```

`vendor`可以参考[这篇博客](http://ipfans.github.io/2016/01/golang-vendor/)

每次引入新的package都要从新make一下

### 启动容器

> fig up -d

之后就可以自由编辑代码了, 每次保存代码之后`gin`都会自动编译运行你的代码

查看log:

> fig logs beego

negroni
httprouter
golang.org/x/net/trace
github.com/gorilla/mux
github.com/jinzhu/gorm"
github.com/gorilla/schema
