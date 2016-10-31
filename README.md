# site-demo

基于`docker`搭建的Golang开发套件

内含基于`httprouter`搭建的路由系统 和热编译工具

# install

使用前需要安装`docker` `godep`

此处不介绍`docker`的安装配置。 `godep`安装：

```
go get github.com/tools/godep
```

当然，使用其他支持`vendor`的包管理工具代替`godep`是完全可以的

将改项目置于`$HOME/go/src/github.com/zykzhang/`目录下，或者根据实际情况修改`docker/docker-compose.yml`文件

# 启动

```
make server     # 启动服务
make log        # 显示goserver日志
```
其他命令请查看`Makefile`

## 说明

因为`golang`的编译器在容器内,容器内只能看到`volume`范围的东西, 所以要用到`vedor`把包依赖都整理起来放到`volume`范围内,
这样容器内的程序才能正常编译运行

## TODO

orm
