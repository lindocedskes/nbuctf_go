



## B. K8s 部署(k3s-master & slave )

参考博客：[1. docker+k3s安装gzctf](https://blog.csdn.net/a142151/article/details/130944500)                [2. k3s 安装](https://blog.baicai.me/article/2023/quick_start/)

## 一、准备工作

### 1. 准备三台服务器（1台也可以，web和k3s-master在同一个服务器内）

| name       | ip                       | 编号   |
| ---------- | ------------------------ | ------ |
| web        | XX.XX.XX.XX              | 2h2g3M |
| k3s-master | 112.124.42.169（举例ip） | 2h2g-1 |
| k3s-slave  | 8.149.129.232            | 2h4g-2 |

### 2. 修改用户名

```shell
hostnamectl set-hostname web  # web服务器执行
hostnamectl set-hostname k3s-master
hostnamectl set-hostname k3s-slave
```

 3. ​	开放端口：6443
    建议关闭ufw
    
    ```
    ufw disable
    ```
    
    
    
    4、关闭swap分区，关闭selinux/AppArmor（k3s-master和k3s-slave机器执行）

    ```
    swapoff -a
    sed -i ' / swap / s/^\(.*\)$/#\1/g' /etc/fstab
    
    # selinux 在ubuntu默认无
    ```
    
    

​	5、（似乎不用）配置hosts解析（k3s-master和k3s-slave机器执行）

```	
cat >>/etc/hosts<<EOF
112.124.42.169 k3s-master
8.149.129.232 k3s-slave
EOF
```

 6. 初始化更新：

    ```shell
    sudo apt update
    sudo apt upgrade -y
    sudo reboot
    sudo apt install -y curl wget
    ```

    

### 二、安装docker（k3s-master和k3s-slave机器执行）

```shell
curl https://releases.rancher.com/install-docker/20.10.sh | sh
 
systemctl enable --now docker  # docker开启自启
```

### 三、安装k3s集群

​	**法一：多云环境下部署**（测试成功）

```shell
# 1. 在 k3s-master（112.124.42.169）上安装 k3s，并设置其为集群的 master 节点。
curl -sfL https://rancher-mirror.rancher.cn/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn  sh -s - --node-external-ip="112.124.42.169" --flannel-backend=wireguard-native --flannel-external-ip


# 2. Agent 节点 k3s-slave，你需要从 master 节点获取 `K3S_TOKEN` 和 `K3S_KUBECONFIG_FILE`。你可以使用以下命令：
K3S_TOKEN=$(sudo cat /var/lib/rancher/k3s/server/node-token)

curl -sfL https://rancher-mirror.rancher.cn/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn K3S_URL=https://112.124.42.169:6443 K3S_TOKEN=K10a967c42b0217e1d7616f028e1c83e4326cebf740d7aeda8daec684bc42dd4695::server:33ab4bcd170da45928dc5728651db9d6 sh -s - --node-external-ip="8.149.129.232"

kubectl get node #server 测试

#重新安装后，重新获取，配置文件，用与后端生成k8s控制器
/etc/rancher/k3s/k3s.yaml -> kube-config.yaml

```

**关键就是成功获得**k3s.yaml文件

​	**法二：同一局域网（没试过）**

​	**a1**、k3s-master节点执行：k3s-master和gzctfweb可在同一服务器

```shell
curl -sfL https://rancher-mirror.rancher.cn/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn sh -
# 检查
systemctl status k3s

# systemctl enable --now k3s  # k3sserver自启

cat /var/lib/rancher/k3s/server/node-token # 查看token
```

​	**a2、**k3s-slave节点执行：（最好同一局域网，K3S_URL用局域网ip）

```shell
# myserver：master节点的ip、mytoken就是在master节点查看到的内容
# /etc/systemd/system/k3s-agent.service
curl -sfL https://rancher-mirror.rancher.cn/k3s/k3s-install.sh | INSTALL_K3S_MIRROR=cn K3S_URL=https://112.124.42.169:6443 K3S_TOKEN=K10162947d31566a455e95877bf69cfd89560daf4ffe32478ea6eca95c6e0824f9a::server:e839ec36e18fc1209c36703888688bf6 sh -

# systemctl enable --now k3s-agent   # 开启自启

curl -vks https://112.124.42.169:6443/ping #测试
```

​	

### 四、k3s集群 配置更改

1. 更改 NodePort 端口范围（k3s-master机器执行），默认为 30000-32767 （选）

```shell
vim /etc/systemd/system/k3s.service

#添加：kube-apiserver-arg后面是 = ,网上和官网教程错！！
--kube-apiserver-arg=service-node-port-range=20000-50000
    
ExecStart=/usr/local/bin/k3s \
    server \
	'--node-external-ip=112.124.42.169' \
	'--flannel-backend=wireguard-native' \
	'--flannel-external-ip' \
	'--kube-apiserver-arg=service-node-port-range=20000-50000' \
	
	

# 重载
sudo systemctl daemon-reload

sudo systemctl restart k3s
```

​	

2.  更改 K3s 的容器数量限制（k3s-master和k3s-slave机器执行）（选）

```shell
vim /etc/rancher/k3s/kubelet.config
# 编写如下内容：
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
maxPods: 500  # 容器数量限制 默认110
```

3. 添加容器镜像仓库（k3s-master机器执行）

```shell
vim /etc/rancher/k3s/registries.yaml
# 写入以下内容：
mirrors:
  "docker.io":
    endpoint:
      - "https://hub.docker.com"

#重启
sudo systemctl restart k3s
```

​	3、重启k3s集群  (选)

```shell
systemctl daemon-reload && systemctl restart k3s  # master节点
systemctl daemon-reload && systemctl restart k3s-agent # slave节点
#检查是否成功
kubectl get nodes -o wide
```



​	卸载重装：Uninstalling K3s

```shell
# server node
/usr/local/bin/k3s-uninstall.sh
# agent node
/usr/local/bin/k3s-agent-uninstall.sh 
```

​	其他常用：

```shell
# 停止 K3s
/usr/local/bin/k3s-killall.sh
# 手动启动 k3s-slave，通过systemd 服务来启动
sudo systemctl start k3s-agent
# 检查服务状态
sudo systemctl status k3s-agent
# 查看报错日志
cat /var/log/syslog
journalctl -xeu k3s-agent.service #agent
journalctl -u k3s --since "1 hour ago" #server

# master 查看所有注册节点,获取集群节点信息
kubectl get nodes -o wide
# 获取集群节点详细信息 k3s-master 是节点名
kubectl describe node k3s-master 
```

5. 查看负载

   ```shell
   #cpu mem
   top
   
   # master 查看所有注册节点,获取集群节点信息
   kubectl get nodes -o wide
   # 获取集群节点详细信息 k3s-master 是节点名
   kubectl describe node k3s-master 
   ```
   


