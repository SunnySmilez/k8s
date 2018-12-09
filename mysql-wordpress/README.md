# 创建mysql-wordpress

1. 创建PersistentVolume
```bash
kubectl create -f local-volumes.yaml
```

2. 查看创建详情
```bash
kubectl get pv
```

3. 创建配置mysql密码的Secret
```bash
kubectl create secret generic mysql-pass --from-literal=password=YOUR_PASSWORD
```

4. 验证secret是否存在
```bash
kubectl get secrets
```

5. 部署MySQL
```bash
kubectl create -f mysql-deployment.yaml
```

6. 验证pod创建
```bash
kubectl get pods
```

7. 部署WordPress
```bash
kubectl create -f wordpress-deployment.yaml
```

8. 查看service是否正常
```bash
kubectl get services wordpress
```

9. 查看WordPress访问路径
```bash
minikube service wordpress --url
```

10. 删除Secret
「以上整体搭建完成，测试例子删除」
```bash
kubectl delete secret mysql-pass
```

11. 删除Services，Deployments
```bash
kubectl delete deployment -l app=wordpress
kubectl delete service -l app=wordpress
```

12. 删除PersistentVolumeClaim和PersistentVolume
```bash
kubectl delete pvc -l app=wordpress
kubectl delete pv local-pv-1 local-pv-2
```

> [mysql-wordpress-persistent-volume](https://k8smeetup.github.io/docs/tutorials/stateful-application/mysql-wordpress-persistent-volume/)