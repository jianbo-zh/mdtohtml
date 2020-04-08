---
Tags: mirror, proxy
---
**docker hub 加速器（registry mirror）**
针对 docker官方提供的仓库，由于国内网络环境导致下载速度慢，就提出了 docker hub 加速器的概念，所谓加速器 registry mirror 其实就是搭建一个registry，然后将docker hub的registry数据缓存到自己本地的registry。
[搭建registry mirror](https://niyanchun.com/deploy-registry-mirror.html)

整个过程是：当我们使用docker pull去拉镜像的时候，会先从我们本地的registry mirror去获取镜像数据，如果不存在，registry mirror会先从docker hub的registry拉取数据进行缓存，再传给我们。而且整个过程是流式的，registry mirror并不会等全部缓存完再给我们传，而且边缓存边给客户端传。

配置方法：
```
cat /etc/docker/daemon.json 
{"registry-mirrors": ["https://registry.docker-cn.com", "http://hub-mirror.c.163.com", "http://f1361db2.m.daocloud.io"], "insecure-registries":["192.168.190.140:5000"]}
```

**HTTP 代理**
针对非Docker官方仓库的镜像，如果国内无法访问的话，可以通过 HTTP代理来间接下载

配置方法：
修改docker的 systemd 服务，添加Environment环境变量，配置HTTP,HTTPS代理（注意没有分行）
/lib/systemd/system/docker.service

[Service]
#注意未分行，双引号之间用空格隔开
Environment="HTTP_PROXY=http://jevon:SmIxMjM0@167.179.95.191:8123/" "HTTPS_PROXY=http://jevon:SmIxMjM0@167.179.95.191:8123/" "NO_PROXY=localhost,127.0.0.1,index.docker.io,registry.docker-cn.com,hub-mirror.c.163.com,f1361db2.m.daocloud.io"