---
Tags: install, mirrors
---
**1，安装需要的包**
```
sudo apt install apt-transport-https ca-certificates software-properties-common curl
```

**2，添加 GPG 密钥，并添加 Docker-ce 软件安装源**
* 官方源 https://download.docker.com/linux/ubuntu/gpg
* 中科大源 https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/gpg
* 阿里源 https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg
* 腾讯源 https://mirrors.cloud.tencent.com/docker-ce/linux/ubuntu/gpg

```
sudo curl -fsSL https://mirrors.cloud.tencent.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://mirrors.cloud.tencent.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
```

**3，更新软件包缓存**
```
sudo apt update
```

**4，安装 Docker-ce**
```
sudo apt install docker-ce
```

**5, 配置国内Docker Hub 镜像**
```
# 通过脚本设置(daocloud 要快点)
curl -sSL https://get.daocloud.io/daotools/set_mirror.sh | sh -s http://f1361db2.m.daocloud.io

# 直接修改文件
vi /etc/docker/daemon.json （不存在则新建）
{
   "registry-mirrors": [
       "http://f1361db2.m.daocloud.io",
       "https://mirror.ccs.tencentyun.com",
	   "https://registry.docker-cn.com",
	   "https://docker.mirrors.ustc.edu.cn"
  ]
}
```

**6，查看是否添加成功**
```
# 观察 Registry Mirrors 字段
sudo docker info
```

**7，开机自启动并启动 Docker**
```
sudo systemctl enable docker
sudo systemctl start docker
```

**8，添加当前用户到 docker 用户组，可以不用每次 sudo 运行 docker**
* docker 安装后默认创建了docker组
* sudo groupadd docker
* 添加用户到组后，用户需要重新登录才可用
```
sudo usermod -aG docker $USER
```

**9，测试运行**
```
docker run hello-world
```

**10，安装 docker-compose**
```
sudo curl -L "https://github.com/docker/compose/releases/download/1.25.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose
```