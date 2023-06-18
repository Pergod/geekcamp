# 模块三作业
1. 代码构建
当前路径下执行`go build`命令
2. Dockerfile编写，具体内容见Dockerfile
3. 执行构建
`docker build -t httpserver .`
4. 查看images
`docker images`
5. 运行容器镜像
`docker run --name httpserver -p 8888:80 http server`
6. 验证结果
`curl localhost:8888/healthz`
7. 推送镜像
`docker push `