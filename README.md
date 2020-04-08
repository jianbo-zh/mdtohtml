# MdToHtml

这是一个把 Markdown 笔记生成静态 HTML 文档的工具。

[Demo网址](http://111.231.134.102/index.html)

**开发背景：**

日常习惯用Markdown来写笔记，但试用了一些Markdown笔记软件后发现存在些问题（如：操作不顺手，展示效果也不美观；markdown笔记导出到本地麻烦；markdown笔记发布到网上麻烦；有些好的但收费），为了用Markdown写笔记顺手，所以就写了这个小工具

**此工具优点：**

1，自动添加笔记之间的跳转链接，点下鼠标就可以快速查看
2，提取笔记元数据（目前只支持：标签），并通过标签来关联多个笔记记录
3，支持本地和网络浏览生成模式，使用网络浏览模式生成的文档可快速发布到网站

**两种生成模式：**

1，本地浏览（内部跳转链接，只适合本地访问）
2，网络浏览（生成的HTML文件放到域名根目录后，可用直接访问）

**使用方法：**

```
#查看帮助
go run ./package/ --help

#编译可执行文件
go build -o mdtohtml ./package/

#运行执行文件，如果提供 --local 选项则生成本地浏览模式，否则生成网络浏览模式
/path/to/mdtohtml --src-dir ./example --dst-dir ./example-html --local
```

**目录结构**：
```
#markdown目录结构：
example/
├── Docker
│   ├── docker-install.md
│   ├── docker-proxy.md
│   └── make-base-image.md
├── Git
│   ├── gitignore.md
│   └── git-install.md
└── K8s
    └── kubectl-auto-completion.md

#生成的HTML目录结构：
example-html/
├── Docker
│   ├── docker-install.html
│   ├── docker-proxy.html
│   ├── index.html
│   ├── make-base-image.html
│   └── tag
│       ├── base-image.html
│       ├── install.html
│       ├── mirror.html
│       ├── mirrors.html
│       ├── proxy.html
│       └── rootfs.html
├── Git
│   ├── gitignore.html
│   ├── git-install.html
│   ├── index.html
│   └── tag
│       ├── config.html
│       ├── gitignore.html
│       ├── gitk.html
│       └── install.html
├── index.html
└── K8s
    ├── index.html
    ├── kubectl-auto-completion.html
    └── tag
        ├── auto-completion.html
        └── kubectl.html
```