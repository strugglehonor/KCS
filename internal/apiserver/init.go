package apiserver

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// 创建client set 连接k8s集群
func NewClientSet() *kubernetes.Clientset {
	var kubeConfig string
	config, err := rest.InClusterConfig()

	if err != nil {
		if envVar := os.Getenv("KUBECONFIG"); len(envVar) > 0 {
			kubeConfig = envVar
		}

		kubeconfigPointer := flag.String("kubeconfig", "~/.kube/config", "kubeconfig file")
		flag.Parse()
		kubeConfig = *kubeconfigPointer
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)

	}

	clientset, err := kubernetes.NewForConfig(config)
	return clientset
	
}

func init() {
	initRouter(&gin.Engine{})

}