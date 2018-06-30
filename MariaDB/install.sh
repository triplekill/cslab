# https://downloads.mariadb.org/mariadb/10.2.15/
sudo apt-get update
wget http://mirrors.neusoft.edu.cn/mariadb//mariadb-10.2.15/bintar-linux-x86_64/mariadb-10.2.15-linux-x86_64.tar.gz
tar zxvf mariadb-10.2.15-linux-x86_64.tar.gz
mkdir ~/mariadb
cp -r mariadb-10.2.15-linux-x86_64/* $HOME/mariadb/
cp mariadb-10.2.15-linux-x86_64/support-files/my-small.cnf $HOME/mariadb/my.cnf
cp mariadb-10.2.15-linux-x86_64/support-files/mysql.server $HOME/mariadb/
$HOME/mariadb/scripts/mysql_install_db --defaults-file="$HOME/mariadb/my.cnf" --user=`whoami` --basedir="$HOME/mariadb" --datadir="$HOME/mariadb/data"

$HOME/mariadb/mysql.server start
$HOME/mariadb/bin/mysqladmin -u root password 'toor'
