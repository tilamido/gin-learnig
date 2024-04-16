项目结构
my-social-network/
├── backend/              # 后端目录
│   ├── api/              # 控制器层，处理请求、返回响应
│   ├── cache/            # redis相关操作
│   ├── config/           # 配置文件，存放数据库配置、服务器配置等
│   ├── dao/              # 数据访问对象，与数据库交互，执行 SQL 查询等
│   ├── middleware/       # 中间件，例如跨域访问、日志记录等
│   ├── models/           # 模型层，定义数据结构和 ORM 模型
│   ├── queue/            # 消息队列，定义消息发布和消费函数
│   ├── routes/           # 路由配置，定义 API 路径与处理函数的映射
│   ├── runtime/          # 运行日志
│   ├── uploads/          # 上传文件
│   ├── Dockerfile        # Docker后端配置
│   ├── go.mod            # 后台包
│   ├── go.sum            # 后台插件
│   ├── main.go           # 后台入口
├── frontend/             # 前端目录（Vue.js）
│   ├── node_modules/     # 存放项目依赖的 npm 包，通过运行 npm install 安装
│   ├── public/           # 存放静态资源，如 HTML 入口文件、图标等，这些资源在构建过程中一般会被复制到输出目录
│   ├── src/              # 存放 Vue.js 项目的源代码
│   ├── Dockerfile        # Docker前端配置
│   ├── jsconfig.json     # 为 VS Code 提供项目的 JS 代码基础配置，比如路径别名等
│   ├── package-lock.json # 锁定安装的 npm 包版本，确保其他开发人员安装相同版本的包
│   ├── package.json      # 定义项目的 npm 依赖、脚本等
│   ├── README.md         # 项目的 README 文件，通常包含项目说明、构建步骤等
│   └── vue.config.js     # Vue CLI 的配置文件，用于自定义 Vue CLI 项目的构建、开发服务器等配置
├──  mysql/               # MySQL初始化脚本，用于创建数据库和表
├──  testtools/           # 产生随机测试数据脚本
├──  web                  # 移动端静态文件
└──  docker-compose.yml   # docker服务部署文件

技术栈
后端: Go (使用Gin框架和GORM)
前端: Vue3.js
数据库: MySQL
缓存: Redis (用于存储点赞数和快速读取)
消息队列: RabbitMQ (用于保证数据一致性)
容器化: Docker

核心功能实现
用户认证:
使用Gin框架实现RESTful API，处理登录、注册、修改密码等请求。
使用session进行用户认证。
存储加密后的密码。
发布朋友圈:
用户可以发布文字、图片等信息。
信息存储在MySQL和后端服务器中。
浏览朋友圈:
用户可以浏览自己和别人的朋友圈。
可以根据点赞数查看排行榜。
点赞功能:
利用Redis实现高性能的点赞功能。
redis保存点赞信息，用RabbitMQ控制异步 与mysql数据库同步 实现数据一致
细节：
用户针对某个 朋友圈点赞或者取消点赞的 写入方式采用 缓存-MQ-SQL
针对查询某个朋友圈点赞数 操作采用 缓存-SQL-MQ-缓存
点赞数可以实时更新，并可以根据点赞数对朋友圈进行排序。
个人内容管理:
用户可以管理自己发布的朋友圈内容，比如删除。
删除用户，点赞不消失，
删除朋友圈，点赞数据库删除，同时同步redis
Docker容器化:
使用Docker容器化后端应用、前端应用、RabbitMQ、MySQL和Redis。
使用Docker Compose来定义和运行多容器Docker应用程序。

部署说明：
Win：
1、下载安装docker desktop：
docker官网：https://www.docker.com/get-started
2、构筑项目：
切换到当前目录命令行运行：docker-compose up -d -- build
linux:
1、下载安装Docker、Docker Compose：
命令行：
sudo apt update
sudo apt install apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt update
#安装Docker CE（社区版）
sudo apt install docker-ce
#验证Docker安装
sudo docker run hello-world
#安装Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
#验证安装
docker-compose --version

改进方向：

技术优化：
models 尽量只操作 index，计算放到 api进行处理

一、开发流程与技术选型
1、本项目使用是前后端分离开发模式，什么是前后端分离的开发模式，除了这种模式常用的还有哪些开发模式，相对来说这种模式有哪些优缺点。
3、本项目使用到了哪些编程语言，为什么选择使用这些语言，这些语言有什么优势。
4、代码的版本控制使用的是git工具，什么是git工具，git工具主要能为我们解决软件开发项目中的哪些问题，跟git类似的工具常用的还有哪些，他们之间有什么区别。
5、在前后端分离的开发模式中如何降低前后端的沟通成本提高开发效率，比如使用接口文档，Mock工具之类。
6、后端使用到了 Redis、RabbitMQ、Docker等技术，这些技术分别都是用解决什么场景问题的，以及他们在解决问题的同时是否又会带来其他问题。
7、用户身份验证使用的是jwt技术，什么是jwt,传统的身份验证还有哪些，相比起来jwt有什么优势适合什么场景下使用，jwt的token过期时间如何处理，单点登录了解下。

二、web安全相关
1、什么是sql注入，本项目中是如何防范sql注入的。
2、什么是XSS攻击，本项目中如何防范的。
3、什么是CSRF攻击，本项目中是如何防范的。
4、什么是DDos攻击，本项目中是如何防范的。

三、应用功能难点攻克
1、本项目在数据安全，防止数据泄露上做了哪些防范。
2、本项目是如何解决防止用户随意发布非法入黄赌毒等内容公布到互联网上的。
3、本项目是如何解决用户在浏览别人发布的内容时不会出现重复浏览，可以适当搞点推荐算法相关的东西。
4、本项目是如何控制用户点赞的有效性，防止恶意刷点赞量的。
5、本项目是如何防止用户借助本平台恶意刷广告等垃圾信息的。
