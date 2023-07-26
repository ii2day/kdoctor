# 安装文档

## 介绍

安装 kdoctor 对集群内外的网络及性能进行检查

## 实施要求

1.一套完整的 k8s 集群（推荐集群版本>= 1.22）

2.已安装 [Helm](https://helm.sh/docs/intro/install/)

## 安装

### 添加 helm 仓库

```shell
helm repo add kdoctor https://kdoctor-io.github.io/kdoctor
helm repo update kdoctor
```

### 安装 kdoctor
kdoctor 可以根据不同的需求进行安装，以下为几个场景的推荐安装方式

#### 1.POC 或 E2E 环境

当POC或E2E的情况下，它可以禁用控制器来收集报告，因此不需要安装strogeClass

以下方法 agent 只将报告打印到控制台
```shell 
helm install kdoctor kdoctor/kdoctor \
    -n kube-system --wait --debug --create-namespace \
    --set feature.enableIPv4=true --set feature.enableIPv6=true \
    --set feature.aggregateReport.enabled=false
```

以下方法 kdoctor-controller 将所有报告收集到本地主机的磁盘上，当 kdoctor-controller 被调度到其他节点时，历史报告将不会被迁移。
```shell 

helm  install kdoctor kdoctor/kdoctor \
    -n kube-system --wait --debug --create-namespace \
    --set feature.enableIPv4=true --set feature.enableIPv6=true \
    --set feature.aggregateReport.enabled=true \
    --set feature.aggregateReport.controller.reportHostPath="/var/run/kdoctor/controller"
```

#### 2.生产环境

以下方法将 kdoctor-controller 的收集报告引导到存储，因此,需要安装storageClass

```shell 

helm  install kdoctor kdoctor/kdoctor \
    -n kdoctor --wait --debug --create-namespace \
    --set feature.enableIPv4=true --set feature.enableIPv6=true \
    --set feature.aggregateReport.enabled=true \
    --set feature.aggregateReport.controller.pvc.enabled=true \
    --set feature.aggregateReport.controller.pvc.storageClass=local \
    --set feature.aggregateReport.controller.pvc.storageRequests="100Mi" \
    --set feature.aggregateReport.controller.pvc.storageLimits="500Mi"
```

#### 3.multus 多网卡环境

如果需要测试 kdoctor-agent pod的所有网卡，则应该使用multus注释对 kdoctor agent 进行注释

```shell 

# 将一下内容替换为 multus 的实际配置
MULTUS_DEFAULT_CNI=kube-system/k8s-pod-network
MULTUS_ADDITIONAL_CNI=kube-system/macvlan

helm install kdoctor kdoctor/kdoctor \
    -n kube-system --wait --debug --create-namespace \
    --set feature.enableIPv4=true --set feature.enableIPv6=true \
    --set feature.aggregateReport.enabled=false \
    --set kdoctorAgent.podAnnotations.v1\.multus-cni\.io/default-network=${MULTUS_DEFAULT_CNI} \
    --set kdoctorAgent.podAnnotations.k8s\.v1\.cni\.cncf\.io/networks=${MULTUS_ADDITIONAL_CNI}
```

## 确认 kdoctor 所有组件正常运行

```shell
kubectl get pod -n kdoctor
NAME                                  READY   STATUS    RESTARTS   AGE
kdoctor-agent-gp5mh                   1/1     Running   0          137m
kdoctor-agent-xkjn4                   1/1     Running   0          137m
kdoctor-controller-686b75d6d7-k4dcq   1/1     Running   0          137m
```