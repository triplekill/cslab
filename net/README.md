# 学习网络
* 切换到当前目录
* docker build -t local/nettools:latest .
* docker run --rm -it --mount type=bind,source=$PWD/,target=/workspace local/nettools:latest bash
* 关闭并删除所有容器  docker stop `docker ps -q` && docker rm `docker ps -a -q`
* 删除image docker image rm local/nettools:latest
