通过文件创建deployments

```bash
kubectl apply -f ngnix_deployment.yaml
```

列出deployment创建的pod

```bash
kubectl get pods -l app=nginx            
```

查出指定pod信息

```bash
kubectl describe pod <pod_name>
```

删除deployment

```bash
kubectl delete deployment nginx-deployment        
```    
