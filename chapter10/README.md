
### 安装Loki-Stack
```shell
helm repo add grafana https://grafana.github.io/helm-charts

helm repo update

helm upgrade --install loki grafana/loki-stack --set grafana.enabled=true,prometheus.enabled=true,prometheus.alertmanager.persistentVolume.enabled=false,prometheus.server.persistentVolume.enabled=false
```
![img.png](img.png)

### 将grafana service type修改为NodePort
```shell
kubectl edit svc loki-grafana -oyaml -n default
```
![img_2.png](img_2.png)

### 获取grafana service的用户名密码
![img_3.png](img_3.png)
将红框所示部分的内容，base64解码即可

### 通过浏览器访问
![img_1.png](img_1.png)
![img_4.png](img_4.png)


### 验证
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: http-server
spec:
  containers:
    - name: http-server
      image: pergod/httpserver:v1
      # Readiness探针：检查容器是否已准备好接收流量
      readinessProbe:
        httpGet:
          path: /healthz
          port: 80
        initialDelaySeconds: 30
        periodSeconds: 5
        successThreshold: 2
      # 命令探活： 在容器中执行指定的命令，并检查其退出状态码来判断容器的健康状态
      livenessProbe:
        exec:
          command:
            - cat
            - /tmp/health
        initialDelaySeconds: 15
        periodSeconds: 30
      # 容器停止的等待时间，Kubernetes 会在发送 SIGTERM 信号给容器后等待一段时间，
      # 以便容器能够处理完正在运行的请求。
      # 如果在等待时间内容器仍未停止，Kubernetes 会发送 SIGKILL 信号来强制终止容器。
      terminationGracePeriodSeconds: 30
```
