docker build -t kubia .
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
