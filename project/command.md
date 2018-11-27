通过文件创建后端deployments

```shell
kubectl apply -f backend_deployment.yaml
```

查看deployment详细信息

```shell
kubectl describe deployments hello
```
列出deployment创建的pod

```shell
kubectl get pods -l app=hello            
```

查出指定pod信息

```shell
kubectl describe pod <pod_name>
```

创建后端service

```shell
kubectl create -f backend_service.yaml 
```

创建前端应用（deployment和service一起创建）
```shell
kubectl create -f front.yaml 
```

查看service信息
```shell
kubectl get service frontend
```

由于镜像内容如下，所以需要fronted.conf文件
```shell
FROM nginx:1.9.14

RUN rm /etc/nginx/conf.d/default.conf
COPY frontend.conf /etc/nginx/conf.d
```

>[使用config Map解耦容器和配置文件](https://k8smeetup.github.io/docs/tasks/configure-pod-container/configmap/)

>[使用service把前端连接到后端](https://k8smeetup.github.io/docs/tasks/access-application-cluster/connecting-frontend-backend/)