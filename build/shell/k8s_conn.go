package shell

import (
	"errors"
	appsv1beat "k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClient struct {
}

func (c *K8sClient) ClientConfig(context string) (*restclient.Config, error) {
	if context == "" {
		return nil, errors.New("必须设置上下文")
	}

	// use the current context in kubeconfig
	kubeconfigpath := "/Users/sunnysmilez/.kube/config"
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigpath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *K8sClient) Connect(config *restclient.Config) (*kubernetes.Clientset, error) {
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

// ==============
type k8sDemo struct{}

func createDeployment(deploymentData *appsv1beat.Deployment) {
	// config
	cluster := "minikube" //集群的名字
	conn := &K8sClient{}
	clientconfig, err := conn.ClientConfig(cluster)
	if err != nil {
		return
	}

	// k8sclient
	k8sClient, err := conn.Connect(clientconfig)
	if err != nil {
		return
	}

	// deployment的操作
	namespace := "default"
	deploymentClient := k8sClient.AppsV1beta1().Deployments(namespace)
	deploymentClient.Create(deploymentData)
}

/**
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: test-server
  labels:
	name: test-server
  annotations:
	name: test-server
spec:
  replicas: 2
  template:
	metadata:
	  labels:
		app:  test-server
	spec:
	  restartPolicy: Always
	  containers:
	  - name: test-server
		image: test-php
		imagePullPolicy: Never
		command: ["php-fpm"]
		ports:
		- containerPort: 9000
*/
func createPhpDeployment() {
	deploymentData := &appsv1beat.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-server",
			Labels: map[string]string{
				"name": "test-server",
			},
			Annotations: map[string]string{
				"name": "test-server",
			},
		},
		Spec: appsv1beat.DeploymentSpec{
			Replicas: int32Ptr(2),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "test-server",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            "test-server",
							Image:           "test-php",
							ImagePullPolicy: apiv1.PullNever,
							Command:         []string{"php-fpm"},
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: 9000,
								},
							},
						},
					},
				},
			},
		},
	}
	createDeployment(deploymentData)
}

/**
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: test-nginx
  labels:
	name: test-nginx
  annotations:
	name: test-nginx
spec:
  replicas: 1
  template:
	metadata:
	  name: test-nginx
	  labels:
		app:  test-nginx
	spec:
	  restartPolicy: Always
	  containers:
	  - name: test-nginx
		image: test-nginx
		imagePullPolicy: Never
		command: ["/bin/sh"]
		args: ["-c", "nginx -g 'daemon off;'"]
		ports:
		  - containerPort: 80
*/
func createNginxDeployment() {
	deploymentData := &appsv1beat.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-nginx",
			Labels: map[string]string{
				"name": "test-nginx",
			},
			Annotations: map[string]string{
				"name": "test-nginx",
			},
		},
		Spec: appsv1beat.DeploymentSpec{
			Replicas: int32Ptr(1),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "test-nginx",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:            "test-nginx",
							Image:           "test-nginx",
							ImagePullPolicy: apiv1.PullNever,
							Command:         []string{"/bin/sh"},
							Args:            []string{"-c", "nginx -g 'daemon off;'"},
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	createDeployment(deploymentData)
}

func createService(serviceData *apiv1.Service) {
	// config
	cluster := "minikube" //集群的名字
	conn := &K8sClient{}
	clientconfig, err := conn.ClientConfig(cluster)
	if err != nil {
		return
	}

	// k8sclient
	k8sClient, err := conn.Connect(clientconfig)
	if err != nil {
		return
	}

	// deployment的操作
	namespace := "default"
	serviceClient := k8sClient.CoreV1().Services(namespace)

	serviceClient.Create(serviceData)
}

/**
apiVersion: v1
kind: Service
metadata:
  name: test-server
  labels:
	name: test-server
spec:
  ports:
  - port: 9000
	protocol: TCP
	targetPort: 9000
  selector:
	app: test-server
*/
func createPhpService() {
	serviceData := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-server",
			Labels: map[string]string{
				"name": "test-server",
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports: []apiv1.ServicePort{
				{
					Port:     9000,
					Protocol: apiv1.ProtocolTCP,
					TargetPort: intstr.IntOrString{
						IntVal: 9000,
					},
				},
			},
			Selector: map[string]string{
				"app": "test-server",
			},
		},
	}
	createService(serviceData)
}

/**
apiVersion: v1
kind: Service
metadata:
  name: test-nginx
spec:
  type: NodePort
  ports:
  - port: 80
	nodePort: 30010
	#argetPort: nginx-port
	#protocol: TCP
  selector:
	app: test-nginx
*/
func createNginxService() {
	serviceData := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-nginx",
		},
		Spec: apiv1.ServiceSpec{
			Type: apiv1.ServiceTypeNodePort,
			Ports: []apiv1.ServicePort{
				{
					Port:     80,
					NodePort: 30010,
				},
			},
			Selector: map[string]string{
				"app": "test-nginx",
			},
		},
	}
	createService(serviceData)
}

func int32Ptr(i int32) *int32 { return &i }

func main() {
	createPhpDeployment()
	createPhpService()
	createNginxDeployment()
	createNginxService()
}
