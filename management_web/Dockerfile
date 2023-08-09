#第一阶段构建
FROM node:16.13.2 as builder
WORKDIR /app/kubemanagement-web
#拷贝源码
COPY . .
#安装依赖
RUN npm config set registry https://registry.npmmirror.com
RUN npm install

# 开始构建
RUN npm run build:prod

# 第二阶段构建
FROM nginx
COPY --from=builder /app/kubemanagement-web/dist/ /usr/share/nginx/html/
COPY --from=builder /app/kubemanagement-web/default.conf.template /etc/nginx/templates/default.conf.template
EXPOSE 80
