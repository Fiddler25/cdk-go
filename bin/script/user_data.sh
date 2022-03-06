#!/bin/bash
yum -y update
amazon-linux-extras install php7.2 -y
yum -y install mysql httpd php-mbstring php-xml gd php-gd

systemctl enable httpd.service
systemctl start httpd.service

wget http://ja.wordpress.org/latest-ja.tar.gz /
tar zxvf latest-ja.tar.gz
cp -r wordpress/* /var/www/html/
chown apache:apache -R /var/www/html