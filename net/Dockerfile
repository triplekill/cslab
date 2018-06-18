# Version: 0.0.1
FROM ubuntu:latest
MAINTAINER Exfly "exfly@outlook.com"

# 刷新apt-list，否则安装软件失败
RUN apt-get update

# traceroute
RUN apt-get update \
	&& apt-get install -y \
	traceroute \
	--no-install-recommends \
	&& rm -rf /var/lib/apt/lists/*

# vim更好用一点
RUN apt-get update \
	&& apt-get install -y \
	vim \
	--no-install-recommends \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR /workspace
