http://mirror.bit.edu.cn/apache/kafka/

```
chrome open https://kafka.apache.org/quickstart#quickstart_multibroker
and download the version you need.
tar -xzf kafka{version}.tgz
cd kafka{version}

cp bin/kafka-server-start.sh bin/kafka-server-start.sh.bak
vim bin/kafka-server-start.sh
change the memory size
bin/zookeeper-server-start.sh config/zookeeper.properties
bin/kafka-server-start.sh config/server.properties
```
