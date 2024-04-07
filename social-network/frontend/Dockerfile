# 第一阶段：使用node基础镜像构建前端静态文件
FROM node:latest as build-stage

# 设置工作目录
WORKDIR /app

# 复制项目文件
COPY ./package*.json ./

# 安装项目依赖
RUN npm install

# 复制其他项目文件
COPY . .

# 构建项目
RUN npm run build

# 第二阶段：使用nginx基础镜像来服务静态文件
FROM nginx:stable-alpine as production-stage

# 从构建阶段复制构建结果到nginx的服务目录
COPY --from=build-stage /app/dist /usr/share/nginx/html

# 暴露80端口
EXPOSE 80

# 启动nginx，使用默认的配置文件
CMD ["nginx", "-g", "daemon off;"]
