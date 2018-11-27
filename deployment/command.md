通过文件创建deployments

```yaml
kubectl apply -f ngnix_deployment.yaml
```

列出deployment创建的pod

```yaml
kubectl get pods -l app=nginx            
```

查出指定pod信息

```yaml
kubectl describe pod <pod_name>
```

删除deployment

```yaml
kubectl delete deployment nginx-deployment        
```    
