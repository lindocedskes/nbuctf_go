package core

import (
	"flag"
	"path/filepath"
	"runtime"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func K8s() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	_, filename, _, _ := runtime.Caller(0) // 获取当前文件的路径
	dir := filepath.Dir(filename)          //当前文件所在的目录
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(dir, "../k8sconfig.yaml"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse() //支持命令行 go run main.go -kubeconfig=/path/to/your/kubeconfig ，默认值../k8sconfig.yaml

	//从 kubeconfig 文件创建一个 Kubernetes 配置。
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, err
	}
	//创建一个 Kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
