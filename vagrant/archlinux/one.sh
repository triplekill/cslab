#!/bin/bash

echo $0 # 脚本path
echo $1 # 第一个参数
echo $# # 参数数量
echo $@ # 所有参数

if [ "${1##*.}" = "tar" ]
then 
    echo This appears to be a tarball.
else
    echo At first glance, this does not appear to be a tarball.
fi
