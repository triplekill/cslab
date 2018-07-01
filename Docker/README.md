# my notebook
[docker notebook](https://github.com/ExFly/CsLearning/blob/master/NoteBookForDevelop/%E5%B7%A5%E5%85%B7/Docker.md)

# install docker in linux
```sh
sudo bash install.sh


curl -fsSL https://get.docker.com/ | sh
sudo service docker start # centos中 启动docker
sudo echo '{"registry-mirrors": ["http://hub-mirror.c.163.com"]}' > /etc/docker/daemon.json
sudo docker run hello-world
```

# install docker in win
* install virtualbox
* download docker toolbox
* install docker toolbox

# kubernetes
配置单机kubernetes环境
[基于Rancher镜像的k8s安装](https://www.cnrancher.com/kubernetes-installation/)
[kubernetes安装](https://github.com/kubernetes/minikube)
