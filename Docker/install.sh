curl -fsSL https://get.docker.com/ | sh
service docker start # centos中 启动docker
echo '{"registry-mirrors": ["http://hub-mirror.c.163.com"]}' > /etc/docker/daemon.json
docker run hello-world
