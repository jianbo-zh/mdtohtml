---
Tags: base-image, rootfs
---
# 构建基础镜像
容器本身是共享宿主操作系统内核，所以容器基础系统镜像包本身就是一个标准的 Linux rootfs + 用户自定义的工具。

构建标准的 Linux rootfs 的方式有很多种方法，Redhat、Debian、SUSE等主流的发行版都有提供相应的工具支持。

大概的流程如下：

构建基础的 rootfs --> 配置基础系统参数 --> 部署用户自定义软件 --> 清理系统 --> 打包为容器镜像 --> 测试镜像 --> 发布仓库

例子：
以 Ubuntu 16.04.01 LTS 版为例，制作一个 Ubuntu 16.04 LTS 的 Docker 基础系统镜像：

```
# 1，安装 Debootstrap
sudo apt install debootstrap

# 2, 通过 Debootstrap 构建 Ubuntu 16.04 LTS 的 rootfs
# 创建 rootfs 存放的位置，如我们把新的 rootfs 存放在 /opt/new_os：
sudo mkdir -p /opt/new_os

# 构建基础 Ubuntu 16.04 LTS 的 rootfs（Debootstrap 工具的参数使用 --help 查看）：

sudo debootstrap --verbose --arch=amd64 xenial /opt/new_os http://mirrors.aliyun.com/ubuntu

# 配置基础系统参数：
# a, 切换到新 rootfs ：
 sudo chroot /opt/new_os /bin/bash

# b, 安装基础包（请根据实际需求安装）：
apt -y update && apt -y upgrade && apt -y install vim locales

# c, 配置系统字符集（根据提示进行）：
dpkg-reconfigure locales

# d, 配置时区：
cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# e, 可选：配置第三方衍生系统版本信息（如：UbuntuKylin）
tee /etc/ubuntukylin-release <<-‘EOF'
DISTRIB_ID=Ubuntu Kylin
DISTRIB_RELEASE=16.04
DISTRIB_CODENAME=xenial
DISTRIB_DESCRIPTION="Ubuntu Kylin 16.04"
EOF

# f, 清理系统：
rm -Rf /tmp/*  && apt clean

# g, 退出当前 rootfs
exit

# 3, 打包并创建 Docker 镜像（前置条件：当前系统已经配置了 Docker 运行时环境）：
sudo tar -C /opt/new_os/ -c . | sudo docker import - new_os

# 4, 测试
sudo docker run new_os  cat /etc/lsb-release
```