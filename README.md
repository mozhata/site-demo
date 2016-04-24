# site-demo

### 这是一个半成品

## 期望

搭建一个基于docker容器的网站demo, 为了方便, 先尝试使用beego,计划使用MySQL存放常用读写数据, 使用MongoDB存储常读少改的数据,使用Elasticsearch做索引
前端...实践下React

## 已实现功能:

- 搭载beego框架的容器正常运行
- 热编译

## TODO

- livereload

  目前热编译用的很不舒服, 有时间实现下livereload...或者看看改下bee工具
- 研究下要不要使用`gocrawl`
- 学一下`MongoDB` = =
- 爬取试验数据存入`MongoDB`一遍后续使用`Elasticsearch`索引

# install

## 说明

因为`golang`的编译器在容器内,容器内只能看到`volume`范围的东西, 所以要用到`godep`把包依赖都整理起来放到`volume`范围内,
这样容器内的程序才能正常编译运行,

`bee`工具好像不支撑`godep`,无法再容器内完成热编译. 所以我弄了个`gin`, 感觉没有`bee`那么好用,勉强凑合了

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

>make save_package

保存依赖包

每次引入新的package都要从新save一下

### 启动容器

> fig up -d

之后就可以自由编辑代码了, 每次保存代码之后`gin`都会自动编译运行你的代码

查看log:

> fig logs beego
