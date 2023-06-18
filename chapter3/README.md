# 模块三作业
1. 代码构建 <br/>
当前路径下执行`go build`命令
2. Dockerfile编写，具体内容见Dockerfile
3. 执行构建<br/>
`docker build -t httpserver .`
4. 查看images<br/>
`docker images`
![img_3.png](img_3.png)
5. 运行容器镜像<br/>
`docker run --name httpserver -p 8888:80 httpserver`
![img_4.png](img_4.png)
6. 查看进程<br/>
`docker ps -a`查看所有进程
![img_5.png](img_5.png)
7. 验证功能是否正常<br/>
通过浏览器访问<br/>
![img_7.png](img_7.png)
或者命令`curl http://127.0.0.1:8888/healthz`
![img_6.png](img_6.png)
8. 查看容器进程PID<br/>
`docker inspect -f {{.State.Pid}} httpserver`
![img_2.png](img_2.png)
9. 进入容器<br/>
以下的PID来源于step 8的结果<br/>
`nsenter -t 19077 -u -i -n -p`
10. 查看IP
`ip addr`
![img.png](img.png)
11. 推送镜像 <br/>
   `docker login`登陆自己的账户 <br/>
   `docker tag httpserver pergod/httpserver:v1`打tag <br/>
   `docker push pergod/httpserver:v1`推送到远程仓库 <br/>
   我的镜像：https://hub.docker.com/repository/docker/pergod/httpserver/general
![img_1.png](img_1.png)