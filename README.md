# Golang Learning

## Kubernetes Plans
1. 写一个控制器，当一个应用启动的时候，保证其他的node上也有其镜像
2. Prometheus 部署(DONE) [kube-prometheus](https://github.com/coreos/kube-prometheus)
3. kubectl plugin [plugins](https://github.com/ahmetb/kubectx)
4. 容器内部访问 k8s api [参考](https://www.jianshu.com/p/b1a723033a3c)
5. k8s 首先看 event 和 ns (DONE)
6. 自定义 crd 开发(WIP)
7. k8s ingress 跨域 [ingress](https://blog.csdn.net/u012375924/article/details/94360425)
8. K8s pod 出外网 — (DONE) 请求先到网关（cni），直接使用 host 的路由出外网
9. Docker 集成 ovs [ovs-docs](https://docs.openvswitch.org/en/latest/intro/install/general/#obtaining-open-vswitch-sources)
10. 学习 loadbalancer 实现 [cloud-provider-openstack](https://github.com/kubernetes/cloud-provider-openstack)
11. Iptables 实现负载均衡
12. API Watch 实现[watch](https://www.jianshu.com/p/1cb577f750f0)
13. 分析 `deployment` 创建流程（TODO）

## Kubernetes Docs
1. [kubernetes 经典网络分析](./doc/network.md)
2. [kubectl exec 实现分析](./doc/kubernetes/kubeexec.md)
3. [kubernetes 集群快速搭建](https://github.com/yingjuncao/kubernetes-ansible)
4. [自定义控制器](./doc/kubernetes/controller.md)
5. [Operator 代码分析](./doc/kubernetes/operator.md)

## OpenStack
1. zun 代码分析 （TODO）
2. kuryr 代码分析（TODO）
3. qinling 代码分析（TODO）

## Golang
1. logrus 和 klog 的使用 (DONE) [官方文档](https://github.com/sirupsen/logrus) | [klog官方文档](https://github.com/kubernetes/klog)
2. 并发 goroutine (DONE)
3. slice 排序 (DONE)
4. channel 实现并发和锁 (DONE), 排他锁 (WIP)
5. go-restful (DONE) [官方文档](https://github.com/emicklei/go-restful)
6. cobra (DONE) [官方文档](https://github.com/spf13/cobra)
7. map 的同时读写 [官方文档](https://golang.org/pkg/sync/#Map)
8. context (WIP) [context](https://mp.weixin.qq.com/s/GpVy1eB5Cz_t-dhVC6BJNw)
9. 协程的调度机制(TODO)
10. interface (DOEN)
11. sync.Mutex (TODO)
12. gRPC (WIP)[官方文档](https://github.com/grpc/grpc-go)
13. yaml (DONE) [官方文档](https://github.com/go-yaml/yaml) | [yaml](https://www.jianshu.com/p/84499381a7da)
14. 匿名函数 (DONE)
15. go-gorm (DONE)[中文文档](http://gorm.book.jasperxu.com/) | [官方文档](https://github.com/go-gorm/gorm)
16. golang 操作 rabbitMq（WIP）
