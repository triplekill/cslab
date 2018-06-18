# 安装
* https://downloads.mariadb.org/mariadb/10.2.15/
tar zxvf mariadb-10.2.15-linux-x86_64.tar.gz
sudo apt-get install libaio1
cp -r mariadb-10.2.15-linux-x86_64/* $HOME/mariadb/
cp /vagrant/env/MariaDB/mariadb-10.2.15-linux-x86_64/support-files/my-small.cnf $HOME/mariadb/my.cnf
cp /vagrant/env/MariaDB/mariadb-10.2.15-linux-x86_64/support-files/mysql.server $HOME/mariadb/

./bin/mysqladmin -u root password 'toor'
/home/ubuntu/mariadb/scripts/mysql_install_db --defaults-file='/home/ubuntu/mariadb/my.cnf' --user=ubuntu --basedir="/home/ubuntu/mariadb" --datadir='/home/ubuntu/mariadb/data'

$HOME/mariadb/mysql.server start

# 使用
* $HOME/mariadb/mysql.server start
