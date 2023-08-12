# KubeClusterManagement

Kubernetes 集群管理平台

# Environment

- 后端
    - Golang 1.20
    - client-go 0.25.2
    - gin 1.8.1
    - docker 24.0.4
    - k8s 1.21.3
    - harbo 2.8.3
- 前端
    - vue-element-admin


# Interface

- Pod 管理接口
    - [x] 命名空间列表接口 (kubectl get namespaces)
    - [x] 创建 (kubectl apply -f xxx.yml)
    - [x] 查看详情/列表 (kubectl describe pod / kubectl get pod -n \<namespace\> -o wide)
    - [x] 编辑更新/升级
    - [x] 删除 
        - (kubectl delete pod \<podname\> -n \<namespace\>)
        - (kubectl get deployment -n \<namespace\>)
        - (kubectl delete deployment \<deployment名\> -n \<namespace\>)
    - [x] pod 新增容忍参数 (tolerations)
    - [x] pod 调度模式 (nodeName, nodeSelector, nodeAffinity)
- NodeScheduling 接口
    - [x] node 列表/详情 (kubectl get nodes / kubectl describe node \<nodename\> )
    - [x] node 标签管理 (kubectl label node \<nodename\> <name=value>)
    - [x] node 污点(taint)管理
    - [x] 查看 node 上所有 pod
- 应用与配置分离接口
    - [x] ConfigMap crud 接口 (kubectl get configmap -n \<namespace\>)
    - [x] Secret crud 接口 (kubectl get secret -n \<namespace\>)
