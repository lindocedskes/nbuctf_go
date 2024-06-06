package system

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lindocedskes/global"
	systemReq "github.com/lindocedskes/model/system/request"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

type K8sService struct{}

// 需求：开启一个容器，通过传入镜像地址和内部端口，返回开启的容器ip和外部端口
// 开启容器  镜像地址如：lin088/w4terctf-2023:web001 内部端口：80
// 1.定义一个 Kubernetes Pod，包含 Docker 镜像和内部端口。
// 2.使用 Kubernetes 客户端创建这个 Pod。
// 3.等待 Pod 状态变为 Running。
// 4.获取 Pod 的 IP 地址和外部端口。
func (k8sService *K8sService) OpenContainer(openImage systemReq.OpenImage, userId uint) (string, int, error) {
	//命名规则：必须由小写字母、数字、'-' 或 '.'
	re := regexp.MustCompile(`[^a-z0-9-.]`)
	imageAddrLabel := re.ReplaceAllString(strings.ToLower(openImage.ImageAddr), "")
	uniqueLabel := fmt.Sprintf("container-%s-%s", imageAddrLabel, strconv.Itoa(int(userId))) // 为每个用户按ImageAddr 生成不同的 Pod 标签
	uniqueName := fmt.Sprintf("n8u-pod-%s-%s", imageAddrLabel, strconv.Itoa(int(userId)))    // 为每个用户按ImageAddr 生成不同的 Pod 名称

	userLabel := fmt.Sprintf("user-%d", userId) // 为每个用户生成不同的标签
	// 检查用户已经开启了多少个 Pod
	pods, err := global.NBUCTF_K8S.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "user=" + userLabel,
	})
	if err != nil {
		return "", 0, err
	}
	if len(pods.Items) >= 3 {
		return "", -1, fmt.Errorf("user has already opened 3 pods")
	}

	// 检查是否已经存在具有相同标签的 Pod,获取 Kubernetes 中名为 "default" 的命名空间中,根据Label筛选的所有 Pod
	pods, err = global.NBUCTF_K8S.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=" + uniqueLabel,
	})
	if err != nil {
		return "", 0, err
	}

	if len(pods.Items) > 0 {
		// 如果已经存在具有相同标签的 Pod，则返回已存在的 Pod 的信息
		pod := pods.Items[0]

		// 检查是否已经存在具有相同标签的 Service
		services, err := global.NBUCTF_K8S.CoreV1().Services("default").List(context.TODO(), metav1.ListOptions{
			LabelSelector: "app=" + uniqueLabel,
		})
		if err != nil {
			return "", 0, err
		}
		if len(services.Items) == 0 {
			global.GVA_LOG.Error("no service found for pod", zap.String("pod", pod.Name))
			return "", 0, fmt.Errorf("no service found for pod %s", pod.Name)
		}

		// 从 Service 中获取外部端口
		service := services.Items[0]
		externalPort := service.Spec.Ports[0].NodePort

		// 获取 Node 的 IP 地址
		nodeIP, err := getNodeIP(pod.Spec.NodeName)
		if err != nil {
			return "", 0, err
		}

		return nodeIP, int(externalPort), nil
	}
	// 定义一个 Kubernetes Pod
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: uniqueName + "-", // 为 Pod 名称添加前缀
			Labels: map[string]string{
				"app":  uniqueLabel,
				"user": userLabel, // 添加用户标签
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "n8u-container",
					Image: openImage.ImageAddr,
					Ports: []v1.ContainerPort{
						{
							ContainerPort: int32(openImage.InnerPort),
						},
					},
					Resources: v1.ResourceRequirements{
						Limits: v1.ResourceList{
							v1.ResourceCPU:    resource.MustParse("1"),   // 限制 CPU 使用量
							v1.ResourceMemory: resource.MustParse("1Gi"), // 限制内存使用量
						},
					},
				},
			},
		},
	}

	// 创建 Pod
	createdPod, err := global.NBUCTF_K8S.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return "", 0, err
	}

	// 等待 Pod 状态变为 Running
	var runningPod *v1.Pod
	for i := 0; i < 30; i++ {
		time.Sleep(10 * time.Second) // 每隔 10 秒检查一次 Pod 状态，最多等待 5 分钟
		runningPod, err = global.NBUCTF_K8S.CoreV1().Pods("default").Get(context.TODO(), createdPod.Name, metav1.GetOptions{})
		if err != nil {
			return "", 0, err
		}
		if runningPod.Status.Phase == v1.PodRunning {
			break
		}
	}

	// 创建 Service
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: createdPod.Name + "-service",
			Labels: map[string]string{ // 为 Service 添加的标签，区分用户的Service
				"app": uniqueLabel,
			},
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app": uniqueLabel,
			},
			Ports: []v1.ServicePort{
				{
					Port:       int32(openImage.InnerPort),
					TargetPort: intstr.FromInt(openImage.InnerPort),
				},
			},
			Type: v1.ServiceTypeNodePort,
		},
	}

	createdService, err := global.NBUCTF_K8S.CoreV1().Services("default").Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		// 如果创建 Service 失败，删除 Pod
		return "", 0, err
	}

	// 获取 Node 的 IP 地址
	nodeIP, err := getNodeIP(runningPod.Spec.NodeName)
	if err != nil {
		return "", 0, err
	}

	// 获取 Service 的外部 IP 地址和端口
	// serviceIP := createdService.Spec.ClusterIP
	externalPort := createdService.Spec.Ports[0].NodePort

	// return serviceIP, int(externalPort), nil
	// 返回 Node 的 IP 地址和外部端口
	return nodeIP, int(externalPort), nil
}

// 获取 Node 的 IP 地址
func getNodeIP(nodeName string) (string, error) {
	// 获取 Node 列表
	node, err := global.NBUCTF_K8S.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	// 获取 Node 的 IP 地址
	var nodeIP string
	if nodeName == "k3s-slave" {
		//Kubernetes API 不提供获取公有 IP 地址的方法，这里硬编码公有云 IP 地址
		// nodeIP = "8.149.129.232" // 硬编码公有云 IP 地址
		nodeIP = global.NBUCTF_CONFIG.K3s.K3sSlave // 改为从配置文件中获取
	} else if nodeName == "k3s-master" {
		//Kubernetes API 不提供获取公有 IP 地址的方法，这里硬编码公有云 IP 地址
		// nodeIP = "112.124.42.169" // 硬编码公有云 IP 地址
		nodeIP = global.NBUCTF_CONFIG.K3s.K3sMaster // 改为从配置文件中获取
	} else {
		for _, addr := range node.Status.Addresses {
			if addr.Type == v1.NodeInternalIP {
				nodeIP = addr.Address
				break
			}
		}
	}

	return nodeIP, nil
}

// 关闭容器
func (k8sService *K8sService) CloseContainer(closeImage systemReq.CloseImage, userId uint) (int, error) {
	//命名规则：必须由小写字母、数字、'-' 或 '.'
	re := regexp.MustCompile(`[^a-z0-9-.]`)
	imageAddrLabel := re.ReplaceAllString(strings.ToLower(closeImage.ImageAddr), "")
	uniqueLabel := fmt.Sprintf("container-%s-%s", imageAddrLabel, strconv.Itoa(int(userId))) // 为每个用户按ImageAddr 生成不同的 Pod 标签

	// 检查是否已经存在具有相同标签的 Pod
	pods, err := global.NBUCTF_K8S.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=" + uniqueLabel,
	})
	if err != nil {
		err = fmt.Errorf("no pod found for user %d", userId) //
		return 0, err
	}

	if len(pods.Items) == 0 {
		return 1, fmt.Errorf("no pod found for user %d", userId)
	}

	// 删除 Pod
	err = global.NBUCTF_K8S.CoreV1().Pods("default").Delete(context.TODO(), pods.Items[0].Name, metav1.DeleteOptions{})
	if err != nil {
		return 0, err
	}

	// 查询 Service 是否存在
	services, err := global.NBUCTF_K8S.CoreV1().Services("default").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=" + uniqueLabel,
	})
	if err != nil {
		return 0, err
	}

	if len(services.Items) == 0 {
		return 2, fmt.Errorf("no service found for user %d", userId)
	}
	// 删除 Service
	err = global.NBUCTF_K8S.CoreV1().Services("default").Delete(context.TODO(), services.Items[0].Name, metav1.DeleteOptions{})
	if err != nil {
		return 0, err
	}

	return 0, nil
}

// 查询容器运行状态，正常返回 return nodeIP, int(externalPort), nil
func (k8sService *K8sService) CheckContainer(checkImage systemReq.CheckImage, userId uint) (string, int, error) {
	//命名规则：必须由小写字母、数字、'-' 或 '.'
	re := regexp.MustCompile(`[^a-z0-9-.]`)
	imageAddrLabel := re.ReplaceAllString(strings.ToLower(checkImage.ImageAddr), "")
	uniqueLabel := fmt.Sprintf("container-%s-%s", imageAddrLabel, strconv.Itoa(int(userId))) // 为每个用户按ImageAddr 生成不同的 Pod 标签

	// 检查是否已经存在具有相同标签的 Pod
	pods, err := global.NBUCTF_K8S.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=" + uniqueLabel,
	})
	if err != nil { //查询异常
		return "", 3, err
	}

	if len(pods.Items) == 0 { //容器不存在
		return "", 1, fmt.Errorf("no pod found for user %d", userId)
	}

	// 检查 Pod 状态
	runningPod := pods.Items[0]
	if runningPod.DeletionTimestamp != nil { //是否正在被删除
		return "", 2, fmt.Errorf("pod %s is terminating", runningPod.Name)
	}
	if runningPod.Status.Phase == v1.PodPending { // 是否正在创建
		return "", 4, fmt.Errorf("pod %s is being created", runningPod.Name)
	}
	fmt.Println("runningPod:", runningPod.Status.Phase)

	// 根据 pod 获取 Node 的 IP 地址 和 Service 的外部端口
	nodeIP, err := getNodeIP(runningPod.Spec.NodeName)
	if err != nil { //查询异常
		return "", 5, err
	}

	// 获取 Service 的外部 IP 地址和端口
	services, err := global.NBUCTF_K8S.CoreV1().Services("default").List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=" + uniqueLabel,
	})
	if err != nil {
		return "", 5, err
	}
	if len(services.Items) == 0 {
		return "", 5, fmt.Errorf("no service found for pod %s", runningPod.Name)
	}
	service := services.Items[0]
	if len(service.Spec.Ports) == 0 {
		return "", 5, fmt.Errorf("no ports found for service %s", service.Name)
	}
	externalPort := service.Spec.Ports[0].NodePort

	// 返回 Node 的 IP 地址和外部端口
	return nodeIP, int(externalPort), nil
}

// 关闭所有容器
func (k8sService *K8sService) CloseAllContainer() (int, error) {

	// 获取所有 Pod
	pods, err := global.NBUCTF_K8S.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return 2, err
	}
	// 获取所有 Service
	services, err := global.NBUCTF_K8S.CoreV1().Services("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return 2, err
	}
	//查询是否还存在 Pod和Service
	fmt.Println("pods:", len(pods.Items))
	fmt.Println("services:", len(services.Items))
	if len(pods.Items) == 0 && len(services.Items) == 1 { //只有一个默认的service
		return 0, nil
	}

	// 删除所有 Pod
	for _, pod := range pods.Items {
		err = global.NBUCTF_K8S.CoreV1().Pods("default").Delete(context.TODO(), pod.Name, metav1.DeleteOptions{})
		if err != nil {
			return 2, err
		}
	}

	// 删除所有 Service
	for _, service := range services.Items {
		err = global.NBUCTF_K8S.CoreV1().Services("default").Delete(context.TODO(), service.Name, metav1.DeleteOptions{})
		if err != nil {
			return 2, err
		}
	}

	return 1, nil
}
