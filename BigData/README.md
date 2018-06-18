https://hadoop.apache.org/docs/r1.0.4/cn/hdfs_user_guide.html
[hadoop](http://hadoop.apache.org/releases.html)
[zookeeper](http://mirror.bit.edu.cn/apache/zookeeper/zookeeper-3.4.12/)
[hbase](http://www.apache.org/dyn/closer.lua/hbase/2.0.0/hbase-2.0.0-bin.tar.gz)
[]()
[]()

# 步骤
* 下载zk hadoop
* tar -xzvf zookeeper
* tar -xzvf hadoop

# 安装redis
sudo apt-get install tcl
make
make PREFIX=/home/ubuntu/redis-4.0.10 install
<!-- make MALLOC=libc -->
# 安装 couchDB
admin admin

# MongoDB
https://www.jianshu.com/p/5598f1dcbb98
## apt-get
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927
echo "deb http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.2.list
sudo apt-get update && sudo apt-get install -y mongodb-org
sudo apt-get install -y mongodb-org=3.2.9 mongodb-org-server=3.2.9 mongodb-org-shell=3.2.9 mongodb-org-mongos=3.2.9 mongodb-org-tools=3.2.9

## 手动
curl -O https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu1604-3.6.5.tgz
tar -zxvf mongodb-linux-x86_64-ubuntu1604-3.6.5.tgz
mkdir -p /data/db
sudo chmod 755 /data/*
mongod --dbpath ${MONGO_PATH}/data

# MySQL
https://www.jianshu.com/p/117dab1b658d
sudo apt install make cmake gcc g++ perl bison libaio-dev libncurses5 libncurses5-dev libnuma-dev
curl -O https://dev.mysql.com/get/Downloads/MySQL-5.7/mysql-5.7.22.tar.gz
tar xzvf mysql-5.7.22.tar.gz
cd mysql-5.7.20
cmake .
tar xzvf boost_1_59_0.tar.gz
cd boost_1_59_0
sudo ./bootstrap.sh
sudo ./b2 install
cd mysql-5.7.17
cmake . -DBUILD_CONFIG=mysql_release -DCPACK_MONOLITHIC_INSTALL=ON -DCMAKE_INSTALL_PREFIX=/usr/local/mysql -DDEFAULT_CHARSET=utf8 -DDEFAULT_COLLATION=utf8_general_ci -DMYSQLX_TCP_PORT=33060 -DMYSQL_UNIX_ADDR=/usr/local/mysql/mysql.sock -DMYSQL_TCP_PORT=3306 -DMYSQLX_UNIX_ADDR=/usr/local/mysql/mysqlx.sock -DMYSQL_DATADIR=/usr/local/mysql/data -DSYSCONFDIR=/usr/local/mysql/etc -DENABLE_DOWNLOADS=ON -DWITH_BOOST=system
sudo make
sudo make install
sudo groupadd mysql
sudo useradd -r -g mysql -s /bin/false mysql
cd /usr/local/mysql
sudo chown -R mysql .
sudo chgrp -R mysql .
sudo bin/mysqld --initialize --user=mysql