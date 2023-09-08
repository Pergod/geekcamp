

### istio安装
```shell
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.12.0
cp bin/istioctl /usr/local/bin
istioctl install --set profile=demo -y

#如果运行Curl命令提示Connection refused，可以访问https://istio.io/downloadIstio,将对应的Bash文件，拷贝下来执行
#具体内容见istio-install.sh
```
安装完成，kubectl get pod -n istio-system
![img.png](img.png)


### 创建namespace（非必要步骤） 
```shell
kubectl create ns sidecar
kubectl label ns sidecar istio-injection=enabled
```

### 发布为ingress gateway(https)
```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: https-service0
  namespace: tracing
spec:
  gateways:
    - https-service0
  hosts:
    - 'cncamp.com'
  http:
    - match:
        - uri:
            exact: /service0
      route:
        - destination:
            host: service0
            port:
              number: 80
---
apiVersion: networking.istio.io/v1beta1
kind: Gateway
metadata:
  name: https-service0
  namespace: tracing
spec:
  selector:
    istio: ingressgateway
  servers:
    - hosts:
        - 'istioexam.com'
      port:
        name: https-service0
        number: 443
        protocol: HTTPS
      tls:
        mode: SIMPLE
        credentialName: istioexam-credential
```