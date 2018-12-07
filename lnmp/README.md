我使用的是mac下的minikube，所以直接minikube ssh 登录虚拟机，创建本地image
创建php容器

```shell
docker build -f  ./php/build/Dockerfile -t test-php .
```

创建nginx容器

```shell
cd nginx/build
docker build -t test-nginx .
```

创建php

```shell
kubectl create -f lamp/php/deployment.yaml 
kubectl create -f lamp/php/serverce.yaml
```

创建nginx

```shell
kubectl create -f lamp/nginx/deployment.yaml
kubectl create -f lamp/nginx/serverce.yaml 
```

注：
先创建php服务，再创建nginx服务
先创建deployment，再创建service