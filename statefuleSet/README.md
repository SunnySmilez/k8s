# statefulSet示例

1. nginx
a. 一个名为 nginx 的 headless service，用于控制网络域
b. 一个名为 web 的 StatefulSet，它的 Spec 中指定在有 3 副本，每个 Pod 中运行一个 nginx 容器
c. volumeClaimTemplates 使用 PersistentVolume Provisioner 提供的 PersistentVolumes 作为稳定存储

> [statefuleSer](https://k8smeetup.github.io/docs/concepts/workloads/controllers/statefulset/)

