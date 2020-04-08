---
Tags: install, config, gitk
---
## 安装
**CentOS**
```
sudo yum install git
```

**Ubuntu**
```
sudo apt-get install git
```

## GITK-图形化工具
gitk 是 git 提供的一个gui工具，可以很清晰地查看搜索提交历史及 git 相关操作。在终端 git 仓库目录下输入 gitk 命令即可使用
按照git后，系统默认安装，如果没有gitk命令，请更新git

## 配置
```
# 设置全局配置-个人信息
git config --global user.name "John Doe"
git config --global user.email johndoe@example.com

# 设置全局配置-编辑器
git config --global core.editor emacs

# 查看配置
git config --list

# 查看单个配置
git config user.name
```

## 获取帮助
```
# git help <verb>
git help config

#man git-<verb>
man git-config
```