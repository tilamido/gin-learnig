FROM mysql:latest

# 将初始化脚本复制到容器中
COPY init.sql /docker-entrypoint-initdb.d/

# 授权执行权限
RUN chmod +x /docker-entrypoint-initdb.d/init.sql
