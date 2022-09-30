[toc]

# 9.23-9.30

## 学习内容

### Golang

- 了解了 Golang 的主要特性和基本语法。
- Go 在本地环境的安装和环境变量配置
- Go 模块（module）和包（package）的设计，模块管理工具（go mod）
- 了解常用的标准库的使用（fmt/runtime/bufio/log/net/os/errors/time/refactor/testing/sync）
- Go 包的单元测试（unit test）和基准测试（benchmark）
- 错误和异常的处理（defer/panic/revocer）
- 熟悉 VS Code 调试技巧
- 了解 Go 的并发编程思想，并发编程（goroutine/sync/channel/select）
- 实现一个类似 Gin 的 Web 框架（`Doing`）

### Linux

- 了解进程和线程相关知识
- - 串行程序、并发程序、并行程序、并发系统/并行系统；
- - 高内聚、松耦合
- - 同步的概念、作用和原则
- - 内部通信、外部通信、通信缓存
- - IPC
- - 程序、进程、父进程、子进程、兄弟进程，进程树，pid、ppid，进程状态和生命周期
- - 系统调用
- - COW、数据段、堆、栈 副本、共享代码段
- - 进程空间：用户空间、内核空间
- - 内存单元地址/虚拟内存地址、指针、内存寻址，物理内存、虚拟内存
- - 内核为每个用户进程分配虚拟内存，内核空间分配给内核专用。
- - 页、页框，共享内存区、swap分区
- - CPU两种状态：内核态、用户态，系统调用函数/内核函数
- - 进程切换的概念和代价，进程调度
- - ...

### 容器云

- 了解容器技术（docker）、容器集群和容器编排技术（k8s）的发展背景。

#### Docker

- 了解容器和虚拟机的区别。
- 容器的本质是单进程，了解实现容器的关键技术（Namespace、Cgroups）。
- 镜像的本质是文件系统（rootfs），了解实现的关键技术（Mount Namespace、chroot（pivot_root）和 UnionFS）和镜像分层的设计思想。
- 掌握编写 Dockerfile 制作镜像（rootfs），使用一些标准的原语构建镜像。
- 掌握镜像和容器的基本操作命令。
- 了解数据卷挂载的实现原理。

#### K8S

- 了解 k8s 的架构：
- - 两种节点：由 Master 和 Node 组成。
- - Master 节点的三个核心组件及功能（ApiServer、Scheduler、Controller）。
- - 集群持久化存储组件（etcd）
- - Node 节点的核心组件（kubelet）。

- 了解 k8s、Pod、Container 的关系：
- - 容器是可以承载一些业务/项目/软件的沙盒环境，提供了打包、分发和运行的功能。
- - Pod 属于容器和 k8s 的一个中间层，代表了一组容器，k8s 不直接调度容器，Pod 是 k8s 的最小调度单元。
- - k8s 通过调度 Pod 来完成业务（容器）的编排和容器化。
- - 了解容器的 OCI 和 CRI。

- 了解 Controller 的使用：
- - 基本概念：k8s 不直接调度 Pod，而是通过 Controller 来管理 Pod，Controller 定义了 Pod 的部署特性，比如有几个副本、在什么样的 Node 上运行等。
- - k8s 底层的几种编排对象（Pod、Job、CronJob）
- - k8s 对外的几种服务对象（Service、Secret、Horizontal Pod Autoscaler）
- - k8s 负责运行容器的几种 Controller（编排对象管理器）：
- - - Deployment：Pod 的多实例管理器；
- - - ReplicaSet：Pod 的多副本管理器；
- - - DaemonSet：用于每个 Node 最多只运行一个 Pod 副本的场景；
- - - StatefuleSet：保证 Pod 的每个副本在整个生命周期中名称是不变的；
- - - Job：用于运行结束就删除的应用，其他都是长期持续运行的。
- - k8s 负责访问容器的 Service：
- - - 定义了外界访问一组特定 Pod 的方式。Service 有自己的 IP 和端口，Service 为 Pod 提供了负载均衡。

- 了解 k8s 的编排方式：
- - 使用 YMAL 文件来定义`运行容器`方式和`访问容器`的方式。

- 本地部署 k8s 集群的方式：
- - 了解 kubeadm 命令的基本使用：
- - - 部署 Master 节点；
- - - 部署容器网络插件；
- - - 部署 Kubernetes Worker；
- - - 部署 Dashboard 可视化插件；
- - - 部署容器存储插件。
- - 使用 minikube 在本地创建和管理集群（单集群）。
- - - 熟悉 minikube 命令；
- - - 熟悉 kubectl 命令；

## 代码&笔记

- [Github](https://github.com/dszkng/study-record)

## 学习心得

>在云计算领域，国外互联网大厂们通过对 Docker 公司的几轮厮杀和围剿之后，已经形成了定局，Docker 技术依然是容器技术的标准和规范，K8S 成为 PaaS 领域构建基础设施最重要的一个工具，CNCF 是他们共建的云原生开源基金会。

>关于容器云和`SRE`，我的理解：
>在中台开发的所有业务侧应用，都是采用的微服务架构，是把传统的单体多模块大应用，按照微服务化的思路去进行了划分和拆分，一个应用就变成了若干个微服务的组合，从微服务的开发、部署、运维和管理的角度看，每个微服务都是根据业务特点按一定的策略去拆分设计的，都是比较独立的一个单元，每个微服务单元除了需要实现各自不同的业务逻辑之外，其实还有一定的共性的，比如微服务本质上都是一个个容器，容器是作为微服务的一个载体和存在形式，比如容器运行需要的资源都是一样的（cpu/内存/存储/网络带宽等），如果不考虑容器的内容（业务逻辑），这么些容器是可以统一进行管理或治理的，这其实就是分而治之理念。

>中台就像是应用/服务的生产基地，将物料（从应用/服务市场引入）进行组装（业务代码开发和引用拓扑管理），创建出来一个应用/服务，形成产物（发布）进行交付（上线）。
>1. 在中台层面看，中台每创建一个应用/服务，都会创建一个或多个微服务，微服务之间的关联关系是由注册中心（分为动态和静态）来维护的，微服务之间是通过约定好的标准通信协议（HTTP或RPC）来实现数据交换（中台提供了sdk方式调用）。对微服务的操作就是调部署或网关层API完成部署、回滚、限流、熔断、监控、告警等等操作。
>2. 从底层/部署层面看，每个应用/服务都是一个容器，对一个应用/服务的操作本质上就是对单个或一批容器的操作，比如部署、伸缩、滚动更新、扩缩容、分配资源以及容器之间的通信，是通过 k8s 编排来实现的。

>每个容器都不是无故存在的，都是因为上游某个业务而产生的，业务越多容器越多，站在业务层以下（底层）的视角，看到的是无数的容器，这么多容器在一起应该就是容器云，k8s 是容器的管理者，在 linux 的世界一切皆文件，在 k8s 的世界里一切服务/组件皆容器，`MSP`侧重服务的治理和中间件的集成和维护，为上层的业务（微服务）提供支撑。而`SRE`应该就是维护底层这一堆容器、组件、节点的稳定，为上层`MSP`或部署提供服务编排和集群管理能力。运维应该就是维护云服务器/云服务资源或自建机房服务器稳定吧。

>![](https://s3.bmp.ovh/imgs/2022/09/30/0bea36a74e5bd3f9.jpg)

## 问题记录

1. 开启10个`goroutine`并发执行，打印的结果都是10，打印结果为什么不是乱的？

```
var x = 0

for i := 0; i < 10; i++ {
	fmt.Println("i1=", fmt.Sprintf("%d", i))

	go func() {
		fmt.Println("i2=", fmt.Sprintf("%d", i))
		x++
	}()
}
```

打印结果：

```
i1= 0
i1= 1
i1= 2
i1= 3
i1= 4
i1= 5
i1= 6
i1= 7
i1= 8
i1= 9
i2= 10
i2= 10
i2= 10
i2= 10
i2= 10
i2= 10
i2= 10
i2= 10
i2= 10
i2= 10
10
```

2. 创建channel，设置的缓冲大小的具体起什么作用？

```
var ch1 = make(chan int)
var ch2 = make(chan int, 10)
```

3. 课程中提到 kubeadm 不能容器化，要在宿主机中部署，因为容器（特殊进程）是用 namespace 隔离进程的，管理不了宿主机其他进程。如果 kubeadm 是部署的一个特殊的容器，不指定 namespace 隔离，也不使用 cgroups 来限制访问宿主机资源，并且是其他进程的父进程，理论上是不是可以管理宿主机上其他资源？

4. 关于有状态的容器，Pod1 如果增加副本了，变成两个 Pod，即 Pod1 和 Pod2，我理解的请求转发是在 Service 对象来控制的（负载均衡），对上层的访问者来说有几个 Pod 是没有区别的，扩容的时候，Pod2 和 Pod1 的容器有那种 fork 出来的父子进程或分离进程的关系吗？如果是在不同 Host 的集群节点上，怎么跨主机 fork？

5. 其他node加入集群是不是意味着，把 master 的 apiserver 组件作为统一的控制入口，那如果管理员在 node 节点也部署了一个 apiserver，因为是在同一个集群，通信应该没问题，这种多入口的情况合理吗？

6. k8s master 的 apiserver、scheduler、controller 等组件的高可用是通过在单个 master 节点内实现的高可用，还是依赖 master 节点的高可用实现高可用？

## Todolist & Plan

- Linux
- - 掌握常用命令
- - 掌握 shell 脚本编写
- - 继续了解操作系统知识

- Golang
- - 深入理解 Go 并发编程模型，并发机制原理。

- K8S
- - 继续学习「深入剖析Kubernetes」课程。

- Work
- - 熟悉 [aries-resource-topology](https://gitlab.oneitfarm.com/aries-io/aries) 项目代码和需求。