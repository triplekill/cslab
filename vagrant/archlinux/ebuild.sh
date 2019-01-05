#!/usr/bin/env bash

# https://www.ibm.com/developerworks/cn/linux/shell/bash/bash-3/index.html
if [ $# -ne 2 ]
then 
    echo "Please specify ebuild file and unpack, compile or all"
    exit 1
fi
source ./ebuild.conf
if [ -z "$DISTDIR" ]
then 
    # set DISTDIR to /usr/src/distfiles if not already set
    DISTDIR=/usr/src/distfiles
fi
export DISTDIR
ebuild_unpack() {
    #make sure we're in the right directory 
    cd ${ORIGDIR}
     
    if [ -d ${WORKDIR} ]
    then    
        rm -rf ${WORKDIR}
    fi
    mkdir ${WORKDIR}
    cd ${WORKDIR}
    if [ ! -e ${DISTDIR}/${A} ]
    then
        echo "${DISTDIR}/${A} does not exist.  Please download first."
        exit 1
    fi
    tar xzf ${DISTDIR}/${A}
    echo "Unpacked ${DISTDIR}/${A}."
    #source is now correctly unpacked
}
user_compile() {
    #we're already in ${SRCDIR}
    if [ -e configure ]
    then
        #run configure script if it exists
        ./configure --prefix=/usr
    fi
        #run make
        make $MAKEOPTS MAKE="make $MAKEOPTS"  
} 
ebuild_compile() {
    if [ ! -d "${SRCDIR}" ]
    then
        echo "${SRCDIR} does not exist -- please unpack first."
        exit 1
    fi
        #make sure we're in the right directory  
    cd ${SRCDIR}
    user_compile
}
export ORIGDIR=`pwd`
export WORKDIR=${ORIGDIR}/work
if [ -e "$1" ]
then 
    source $1
else
    echo "Ebuild file $1 not found."
    exit 1
fi
export SRCDIR=${WORKDIR}/${P}
case "${2}" in
    unpack)
        ebuild_unpack
        ;;
    compile)
        ebuild_compile
        ;;
    all)
        ebuild_unpack
        ebuild_compile
        ;;
    *)
        echo "Please specify unpack, compile or all as the second arg"
        exit 1
        ;;
esac
# 接下来：
# 如果在 "$DISTDIR" 没找到源代码，则自动下载
# 通过使用 MD5 消息摘要，验证源代码没有受到破坏
# 如果请求，则将编译过的应用程序安装到正在使用的文件系统，并记录所有安装的文件，以便日后可以方便地将包卸载。
# 如果请求，则将编译过的应用程序打包成 tar 压缩包（以您希望的形式压缩），以便以后可以在另一台计算机上，或者在基于 CD 的安装过程中（如果在构建发行版 CD）安装。