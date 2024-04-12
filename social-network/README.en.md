# social-network

#### Description
核心功能实现
用户认证、发布朋友圈、浏览朋友圈、点赞功能、个人内容管理

#### 用法
Win：
1、下载安装docker desktop：
docker官网：https://www.docker.com/get-started
2、构筑项目：
切换到当前目录命令行运行：docker-compose up -d -- build
常见错误，未安装go、node.js 、vue
##############
安装go：设置代理
安装node.js：没雷
安装vue：雷很多
（1）win10需要开放脚本权限
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned
（2）国内用镜像 
npm install -g @vue/cli
 (3) 先 npm install 再 npm run serve
##############
3、查看项目
前端：localhost:80
后端：localhost:8888
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
2、同上
