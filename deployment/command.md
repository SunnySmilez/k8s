通过文件创建deployments

```shell
kubectl apply -f ngnix_deployment.yaml
```

列出deployment创建的pod

```shell
kubectl get pods -l app=nginx            
```

查出指定pod信息

```shell
kubectl describe pod <pod_name>
```

删除deployment

```shell
kubectl delete deployment nginx-deployment        
```    
