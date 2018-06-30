http://mirror.bit.edu.cn/apache/kafka/

```
wget http://mirror.bit.edu.cn/apache/kafka/1.1.0/kafka_2.11-1.1.0.tgz
tar -xzf kafka_2.11-1.1.0.tgz
cd kafka_2.11-1.1.0

cp bin/kafka-server-start.sh bin/kafka-server-start.sh.bak
vim bin/kafka-server-start.sh
change the memory size
bin/zookeeper-server-start.sh config/zookeeper.properties
bin/kafka-server-start.sh config/server.properties
```
