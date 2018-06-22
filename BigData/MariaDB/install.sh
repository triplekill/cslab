# https://downloads.mariadb.org/mariadb/10.2.15/
# tar zxvf mariadb-10.2.15-linux-x86_64.tar.gz
sudo apt-get update
sudo apt-get install libaio1 libaio-dev libssl-dev build-essential automake autoconf libtool cmake bison libncurses5-dev libjemalloc-dev libaio1 zlibc libidn11-dev libidn11 liblz4
mkdir ~/mariadb
cp -r /vagrant/env/MariaDB/mariadb-10.2.15-linux-x86_64/* $HOME/mariadb/
cp /vagrant/env/MariaDB/mariadb-10.2.15-linux-x86_64/support-files/my-small.cnf $HOME/mariadb/my.cnf
cp /vagrant/env/MariaDB/mariadb-10.2.15-linux-x86_64/support-files/mysql.server $HOME/mariadb/
$HOME/mariadb/scripts/mysql_install_db --defaults-file='/home/ubuntu/mariadb/my.cnf' --user=ubuntu --basedir="/home/ubuntu/mariadb" --datadir='/home/ubuntu/mariadb/data'
$HOME/mariadb/mysql.server start
$HOME/mariadb/bin/mysqladmin -u root password 'toor'
