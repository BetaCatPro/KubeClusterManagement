package initiallize

import (
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"kubemanagement.com/global"
)

// X.509证书双向认证
func K8S() {
	kubeconfig := ".kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	global.KubeConfigSet = clientset
}

// 集群外基于 token 单向认证
func K8SWithToken() {
	// k8s ca证书： /etc/kubernetes/pki/ca.crt
	CAData, err := os.ReadFile("k8s/identity/ca.crt")
	if err != nil {
		panic(err)
	}
	config := rest.Config{
		Host: "https://192.168.65.130:6443",
		// 这个token是用过base64解码得来
		// kubectl get secret kubemanagement-system-token-csw4p -n kubemanagement-system -o jsonpath="{.data.token}" | base64 -d > token
		BearerTokenFile: "k8s_use/identity/token",
		TLSClientConfig: rest.TLSClientConfig{
			CAData: CAData,
		},
	}

	clientset, err := kubernetes.NewForConfig(&config)
	if err != nil {
		panic(err.Error())
	}
	global.KubeConfigSet = clientset
}

func isInCluster() (isInCluster bool) {
	tokenFile := "/var/run/secrets/kubernetes.io/serviceaccount/token"
	_, err := os.Stat(tokenFile)
	if err == nil {
		isInCluster = true
	}
	return
}

// 集群内单向认证
func K8SWithDiscovery() {
	if isInCluster() {
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err)
		}
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		global.KubeConfigSet = clientset
	} else {
		K8S()
	}

}
