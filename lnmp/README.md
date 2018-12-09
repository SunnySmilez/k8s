# 搭建环境
「mac+minikube」
1. mac10.13.6
2. minikube version:v0.30.0
3. kubectl version
    
    
    Client Version: version.Info{Major:"1", Minor:"10", GitVersion:"v1.10.3", GitCommit:"2bba0127d85d5a46ab4b778548be28623b32d0b0", GitTreeState:"clean", BuildDate:"2018-05-21T09:17:39Z", GoVersion:"go1.9.3", Compiler:"gc", Platform:"darwin/amd64"}

# 搭建架构图

[项目搭建源码](https://github.com/SunnySmilez/k8s/tree/master/lnmp)

![k8s搭建lnmp服务架构图](/img/k8s/lnmp.png "lnmp")

在搭建之前先了解几个知识点：
- pod节点通过service暴露端口提供外部访问
可以设置type="NodePort"，访问：http://nodeip:nodePort

「nodeip可以使用minikube ip获得；nodePort在service中指定」

[service type文档](https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types)

- 内部服务之间通信，使用服务名称即可，最终转换成cluster ip

「服务名称在service中设置；cluster ip 可以使用kubectl get service获取」

- service通过selector关联到对应的pod

- deployment中可以设置image地址来指定镜像

# 搭建步骤

- 创建php镜像
```shell
docker build -f  ./php/build/Dockerfile -t test-php .
```

- 创建nginx镜像
```shell
cd nginx/build
docker build -t test-nginx .
```
- 创建php服务
```shell
kubectl create -f lnmp/php/deployment.yaml 
kubectl create -f lnmp/php/serverce.yaml
```

- 创建nginx服务
```shell
kubectl create -f lnmp/nginx/deployment.yaml
kubectl create -f lnmp/nginx/serverce.yaml 
```

- 测试服务连通性

获取本地ip

```shell
minikube ip
```

编辑/etc/hosts文件

```shell
test.com minikube输出的ip值
```

访问输出hello world

```shell
export NODE_PORT=$(kubectl get services/test-nginx  -o go-template='{{(index .spec.ports 0).nodePort}}')
echo NODE_PORT=$NODE_PORT
curl test.com:$NODE_PORT
```

# 遇到的问题
### pod启动问题
- 定位pod问题

查看pod日志
```shell
kubectl pod {pod name}
```

查看pod的event日志
```shell
kubectl describe pods {pod-name}
```

- 镜像拉取失败
登录到虚拟主机，拉取镜像进行测试

```shell
minikube ssh
docker pull image_name
```

- 实例化镜像失败
登录到虚拟主机，创建容器，查看具体信息

```shell
minikube ssh
docker run -itd --name=test test_nginx /bin/bash
```

### 外部访问不通
    
- 确认访问地址正确，上图3位置

查看访问日志，确认请求达到

```shell
kubectl exec -it {pod name} /bin/bash
```

- 确认service，deployment关联正确，上图2位置

确认service能选择到pod
```shell
kubectl get pod --selector="app=test-nginx"
```

- 确认nginx和php服务通信正常，上图4位置
    
    登录nginx的pod节点，ping service_name
    
    如果不通，对比ping显示的ip和php服务的service对应的cluster ip
    
    如果对应不上，重新走一遍创建整个服务流程
    
- 调试过程中，需要安装一些软件
```shell
apt-get update
apt-get install -y {soft name}
```
