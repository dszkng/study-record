
# 安装

```
$ curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64
sudo install minikube-darwin-amd64 /usr/local/bin/minikube
```

# 集群管理

## 启动集群

```
$ minikube start --extra-config=kubelet.cgroup-driver=systemd
```

## 查看节点

>`kubectl`是一个用来跟`k8s`集群进行交互的命令行工具

查看节点列表。

```
$ kubectl get node
NAME       STATUS   ROLES    AGE    VERSION
minikube   Ready    <none>   4m4s   v1.25.0
```

查看节点对象详细信息、状态、事件等。

```
$ kubectl describe node minikube
Name:               minikube
Roles:              <none>
Labels:             beta.kubernetes.io/arch=amd64
                    beta.kubernetes.io/os=linux
                    kubernetes.io/arch=amd64
                    kubernetes.io/hostname=minikube
                    kubernetes.io/os=linux
Annotations:        node.alpha.kubernetes.io/ttl: 0
                    volumes.kubernetes.io/controller-managed-attach-detach: true
CreationTimestamp:  Tue, 27 Sep 2022 15:48:11 +0800
Taints:             <none>
Unschedulable:      false
Lease:
  HolderIdentity:  minikube
  AcquireTime:     <unset>
  RenewTime:       Tue, 27 Sep 2022 16:00:52 +0800
Conditions:
  Type             Status  LastHeartbeatTime                 LastTransitionTime                Reason                       Message
  ----             ------  -----------------                 ------------------                ------                       -------
  MemoryPressure   False   Tue, 27 Sep 2022 16:00:52 +0800   Tue, 27 Sep 2022 15:48:11 +0800   KubeletHasSufficientMemory   kubelet has sufficient memory available
  DiskPressure     False   Tue, 27 Sep 2022 16:00:52 +0800   Tue, 27 Sep 2022 15:48:11 +0800   KubeletHasNoDiskPressure     kubelet has no disk pressure
  PIDPressure      False   Tue, 27 Sep 2022 16:00:52 +0800   Tue, 27 Sep 2022 15:48:11 +0800   KubeletHasSufficientPID      kubelet has sufficient PID available
  Ready            True    Tue, 27 Sep 2022 16:00:52 +0800   Tue, 27 Sep 2022 15:48:11 +0800   KubeletReady                 kubelet is posting ready status
Addresses:
  InternalIP:  192.168.49.2
  Hostname:    minikube
Capacity:
  cpu:                4
  ephemeral-storage:  61202244Ki
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  memory:             8151360Ki
  pods:               110
Allocatable:
  cpu:                4
  ephemeral-storage:  61202244Ki
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  memory:             8151360Ki
  pods:               110
System Info:
  Machine ID:                 40026c506a9948afaae3b0fda0a61c83
  System UUID:                4bdb42d9-1f62-41b9-ac4f-271285a5dfb1
  Boot ID:                    4f5e0209-11af-4e30-afb1-434d1c34813f
  Kernel Version:             5.10.124-linuxkit
  OS Image:                   Ubuntu 20.04.5 LTS
  Operating System:           linux
  Architecture:               amd64
  Container Runtime Version:  docker://20.10.17
  Kubelet Version:            v1.25.0
  Kube-Proxy Version:         v1.25.0
PodCIDR:                      10.244.0.0/24
PodCIDRs:                     10.244.0.0/24
Non-terminated Pods:          (7 in total)
  Namespace                   Name                                CPU Requests  CPU Limits  Memory Requests  Memory Limits  Age
  ---------                   ----                                ------------  ----------  ---------------  -------------  ---
  kube-system                 coredns-565d847f94-wmjkx            100m (2%)     0 (0%)      70Mi (0%)        170Mi (2%)     10m
  kube-system                 etcd-minikube                       100m (2%)     0 (0%)      100Mi (1%)       0 (0%)         11m
  kube-system                 kube-apiserver-minikube             250m (6%)     0 (0%)      0 (0%)           0 (0%)         11m
  kube-system                 kube-controller-manager-minikube    200m (5%)     0 (0%)      0 (0%)           0 (0%)         11m
  kube-system                 kube-proxy-trpgg                    0 (0%)        0 (0%)      0 (0%)           0 (0%)         10m
  kube-system                 kube-scheduler-minikube             100m (2%)     0 (0%)      0 (0%)           0 (0%)         11m
  kube-system                 storage-provisioner                 0 (0%)        0 (0%)      0 (0%)           0 (0%)         10m
Allocated resources:
  (Total limits may be over 100 percent, i.e., overcommitted.)
  Resource           Requests    Limits
  --------           --------    ------
  cpu                750m (18%)  0 (0%)
  memory             170Mi (2%)  170Mi (2%)
  ephemeral-storage  0 (0%)      0 (0%)
  hugepages-1Gi      0 (0%)      0 (0%)
  hugepages-2Mi      0 (0%)      0 (0%)
Events:
  Type    Reason                   Age                From             Message
  ----    ------                   ----               ----             -------
  Normal  Starting                 10m                kube-proxy       
  Normal  NodeHasSufficientMemory  20m (x8 over 20m)  kubelet          Node minikube status is now: NodeHasSufficientMemory
  Normal  NodeHasNoDiskPressure    20m (x8 over 20m)  kubelet          Node minikube status is now: NodeHasNoDiskPressure
  Normal  NodeHasSufficientPID     20m (x8 over 20m)  kubelet          Node minikube status is now: NodeHasSufficientPID
  Normal  NodeAllocatableEnforced  20m                kubelet          Updated Node Allocatable limit across pods
  Normal  RegisteredNode           12m                node-controller  Node minikube event: Registered Node minikube in Controller
  Normal  Starting                 10m                kubelet          Starting kubelet.
  Normal  NodeHasSufficientMemory  10m (x8 over 10m)  kubelet          Node minikube status is now: NodeHasSufficientMemory
  Normal  NodeHasNoDiskPressure    10m (x8 over 10m)  kubelet          Node minikube status is now: NodeHasNoDiskPressure
  Normal  NodeHasSufficientPID     10m (x7 over 10m)  kubelet          Node minikube status is now: NodeHasSufficientPID
  Normal  NodeAllocatableEnforced  10m                kubelet          Updated Node Allocatable limit across pods
  Normal  RegisteredNode           10m                node-controller  Node minikube event: Registered Node minikube in Controller
```

查看指定命名空间下的`pod`列表。

```
$ kubectl get pods -n kube-system 
NAME                               READY   STATUS    RESTARTS      AGE
coredns-565d847f94-wmjkx           1/1     Running   0             12m
etcd-minikube                      1/1     Running   2 (12m ago)   13m
kube-apiserver-minikube            1/1     Running   2 (12m ago)   13m
kube-controller-manager-minikube   1/1     Running   3             13m
kube-proxy-trpgg                   1/1     Running   0             12m
kube-scheduler-minikube            1/1     Running   2 (12m ago)   13m
storage-provisioner                1/1     Running   1 (11m ago)   12m
```
## 停止集群

```
minikube stop
```

## 清空集群

```
minikube delete --all
```

## 安装集群可视化 Web UI 控制台

```
minikube dashboard
```
