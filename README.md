# mydocker

使用namesapce, cgroups等linux特性对**进程**进行访问隔离和资源限制，创建一个自定义容器。

它是一个轻量级的Docker，仅仅使用一些Linux特性，而没有利用[containers](https://containerd.io/)或[runc](https://github.com/opencontainers/runc)。

## 特性
- Control Groups：对资源进行限制（CPU, Memory, Swap, Pids）
- Namespace: 对全局系统资源进行隔离（Network, Mount, UST, PID, IPS）
- Union File System: 将镜像层和读写层进行分层（OveralayS）

## 编译
- 直接使用
`go run main.go run alpine`

- 打包使用
`go build -o mydocker`
`./mydocker`

## 使用
使用：
    mydocker [command]
命令：
    - exec: 运行一个指令在一个已经存在的容器中
    - help: 查看帮助
    - images: 查看本地镜像
    - ps: 查看当前容器
    - run:  运行一条指令在一个新容器中

## 例子
1. 在`alpine:latest`中运行`/bin/sh`:
```bash
mydocker run alpine /bin/sh
mydocker run alpine # 效果同上，因为默认的命令是/bin/sh
```
2. 查看当前所有容器 
```bash
mydocker ps
```
输出：
```
CONTAINER ID            IMAGE                   COMMAND
42645fdba523            b747534ae29d            "/bin/sh"
5729e6b2b1a2            b747534ae29d            "/bin/sh"
```

3. 查看所有本地镜像
```bash
mydocker images
```
输出：
```
REPOSITORY                      TAG             IMAGE ID
library/alpine                  latest          e7d88de73db3
library/mysql                   latest          e9c9e3680bba
library/nginx                   latest          bb129a712c24
library/python                  latest          cd150f52893e
library/redis                   latest          03f00cd78924
library/ubuntu                  latest          7c9c7fed23de
```
