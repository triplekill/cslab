# service

## minikube docker register
* [Sharing a local registry with minikube](https://blog.hasura.io/sharing-a-local-registry-for-minikube-37c7240d0615)

# cmd
* `minikube start --registry-mirror=https://registry.docker-cn.com`

# helm
## usage 
* helm repo update


## [install](https://docs.helm.sh/using_helm/#initialize-helm-and-install-tiller)
* https://www.jianshu.com/p/699c5ced3f87
* `brew install kubernetes-helm`
* `helm version`
* `helm init --upgrade -i registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.12.3 --stable-repo-url https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts`

## error
kubectl create serviceaccount --namespace kube-system tiller

kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller

kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'