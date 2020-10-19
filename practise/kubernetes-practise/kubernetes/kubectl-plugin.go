package main

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 学习实践 kubectl 的 plugin https://github.com/kubernetes/sample-cli-plugin
// https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/sample-cli-plugin

// 编辑可执行文件，编译时 以 kubeclt-xxx开始，然后把二进制拷贝到PATH下，使用kubectl xxx 调用
// https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/

func main() {

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", "config/admin.conf")
	if err != nil {
		panic(err)
	}
	clientset, cErr := kubernetes.NewForConfig(kubeConfig)
	if cErr != nil {
		panic(cErr)
	}
	pods, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
    if err != nil {
        panic(err)
	}
    fmt.Println(pods)
}

