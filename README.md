# Golang learning

## Kubernetes
1. 写一个控制器，当一个应用启动的时候，保证其他的node上也有其镜像
2. Prometheus 部署 [kube-prometheus](https://github.com/coreos/kube-prometheus)
3. kubectl plugin [plugins](https://github.com/ahmetb/kubectx)
1. 容器内部访问 k8s api [参考](https://www.jianshu.com/p/b1a723033a3c)
2.  k8s 首先看 event 和 ns
3.  自定义 crd 开发
4.  k8s ingress 跨域 [ingress](https://blog.csdn.net/u012375924/article/details/94360425)
5.  K8s pod 出外网
6.  Docker 集成 ovs [ovs-docs](https://docs.openvswitch.org/en/latest/intro/install/general/#obtaining-open-vswitch-sources)
7.  学习 loadbalancer 实现 [cloud-provider-openstack](https://github.com/kubernetes/cloud-provider-openstack)
8.  Iptables 实现负载均衡

## Golang
1. logrus 的使用
2. 并发 goroutine —— ok
3. slice 排序 —— ok
4. channel 实现并发和锁, 排他锁 —— ok
5. go-restful
6. cobra

## Docs
1. [网络分析](./doc/network.md)
2. [kubectl exec实现分析](./doc/kube-exec.md)
3. [kubernetes 集群快速搭建](https://github.com/yingjuncao/kubernetes-ansible)
