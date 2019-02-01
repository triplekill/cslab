minikube start --registry-mirror=https://registry.docker-cn.com

docker build -t kubia .
docker build -t vieux/apache:2.0 .
docker run --name kubia-container -p 8080:8080 -d kubia
docker exec -it kubia-container bash
pa aux
docker stop kubia-container
docker rm kubia-container

docker tag kubia luksa/kubia
docker push luksa/kubia
docker run -p 8080:8080 -d luksa/kubia

kubectl describe nodes minikube

kubectl run kubia --image=luksa/kubia --port=8080 --generator=run/v1

kubectl get pods
kubectl get nodes
kubectl get svc
kubectl describe pods pods-name
kubectl expose rc kubia --type=LoadBalancer --name kubia-http

k get replicationcontrollers
kubectl scale rc kubia --replicas=3

minikube dashboard
kubectl explain pods

kubectl create -f kubia-manual.yaml

kubectl get po kubia-manual -o yaml

k port-forward kubia-manual 8888:8080
kubectl get po --show-labels
kubectl get po -L creation_method,env
kubectl get po -l creation_method=manual
kubectl annotate pod kubia-manual mycompany.com/someannotation="foo bar" 

k get pods --namespace kube-system
k get pods -n kube-system

kubectl create namespace custom-namespace
kubectl create -f kubia-manual.yaml -n custom-namespace

k delete pods kubia-manual
kubectl delete po -l creation_method=manual
kubectl delete ns custom-namespace
kubectl delete po --all # del all pods in current ns
kubectl delete all --all

kubectl logs mypod --previous
k edit rc kubia
kubectl scale rc kubia --replicas=10
kubectl delete rc kubia --cascade=false

docker build -t luksa/fortune .

kubectl describe po kube-addon-manager-minikube --namespace kube-system

kubectl get po -o wide

k get pv
k get pvc

k get cm
kubectl create configmap fortune-config --from-file=configmap-files

kubectl get secrets

openssl genrsa -out https.key 2048
openssl req -new -x509 -key https.key -out https.cert -days 3650 -subj /CN=www.kubia-example.com

kubectl create secret generic fortune-https --from-file=fortune-https

exec fortune-https -c web-server -- mount | grep certs

kubectl create secret docker-registry mydockerhubsecret \
  --docker-username=myusername --docker-password=mypassword \
  --docker-email=my.email@provider.com

k proxy 
curl localhost:8001
curl localhost:8001/apis/batch/v1/jobs

minikube start --extra-config=apiserver.Features.Enable-SwaggerUI=true

minikube start --registry-mirror=https://registry.docker-cn.com --extra-config=apiserver.Features.Enable-SwaggerUI=true

kubectl rolling-update kubia-v1 kubia-v2 --image=luksa/kubia:v2
kubectl rolling-update kubia-v1 kubia-v2 --image=luksa/kubia:v2 --v 6

k delete rc --all
k create -f kubia-deployment-v1.yaml --record
k rollout status deployment kubia
kubectl patch deployment kubia -p '{"spec": {"minReadySeconds": 10}}'

while true; do curl http://130.211.109.222; done
k set image deployment kubia nodejs=luksa/kubia:v2
k apply -f kubia-deployment-v2.yaml --record
kubectl rollout undo deployment kubia
kubectl rollout undo deployment kubia --record --to-revision=1
kubectl rollout history deployment kubia
kubectl rollout pause deployment kubia

k proxy
curl localhost:8001/api/v1/namespaces/default/pods/kubia-0/proxy/

curl localhost:8001/api/v1/namespaces/default/services/kubia-public/proxy/

kubectl run -it srvlookup --image=tutum/dnsutils --rm
dig SRV kubia.default.svc.cluster.local

kubectl get componentstatuses
kubectl get po -o custom-columns=POD:metadata.name,NODE:spec.nodeName --sort-by spec.nodeName -n kube-system

k exec etcd-minikube --namespace kube-system etcdctl
kubectl get pods --watch
kubectl get events --watch

export no_proxy=$no_proxy,$(minikube ip)

docker save k8s.gcr.io/fluentd-elasticsearch:v2.2.0 | ( eval $(minikube docker-env) && docker load )

docker save k8s.gcr.io/elasticsearch:v6.2.5 | ( eval $(minikube docker-env) && docker load )
