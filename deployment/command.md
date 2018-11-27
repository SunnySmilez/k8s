通过文件创建deployments
    
    kubectl apply -f ngnix_deployment.yaml
    
列出deployment创建的pod

    kubectl get pods -l app=nginx            
    
查出指定pod信息
    
    kubectl describe pod <pod_name>
    
删除deployment
    
    kubectl delete deployment nginx-deployment        