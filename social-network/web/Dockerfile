# 使用官方的 Nginx 镜像作为基础镜像
FROM nginx:alpine

# 将静态文件复制到 Nginx 期望的服务目录下
COPY . /usr/share/nginx/html
# 将替换后端地址的脚本复制过去
COPY init-script.sh /docker-entrypoint.d/
# 给启动脚本赋予执行权限
RUN chmod +x /docker-entrypoint.d/init-script.sh

# 暴露端口 80，使得外部可以访问你的容器
EXPOSE 80

# 使用 Nginx 默认的启动命令
CMD ["nginx", "-g", "daemon off;"]
