apt-get install libcloog-ppl-dev libcloog-ppl0

tar xzvf drizzle7-2011.07.21.tar.gz
cd drizzle7-2011.07.21/
./configure --without-server
make libdrizzle-1.0
make install-libdrizzle-1.0
======================================
dpkg-reconfigure tzdata
echo 'export TZ="EET"' >> /etc/default/rsyslog
restart rsyslog


add-apt-repository ppa:chris-lea/redis-server
apt-get install redis-server

apt-get install build-essential

apt-get install sqlite3 libsqlite3-dev

apt-get install libpcre3-dev libssl-dev

apt-get install git-core

git clone https://github.com/agentzh/echo-nginx-module.git

git clone https://github.com/agentzh/redis2-nginx-module.git

apt-get install inotify-tools 

apt-get install libgd2-noxpm
apt-get install libgd2-noxpm-dev # for mage_filter_module
  


cp -a nginx/  nginxlast

apt-get autoremove --purge nginx-extras
apt-get autoremove --purge nginx-common

adduser --system --no-create-home --disabled-login --disabled-password --group nginx

./configure --prefix=/opt/nginx --user=nginx --group=nginx --add-module=../echo --add-module=../redis2-nginx-module

./configure --prefix=/opt/nginx --user=nginx --group=nginx --with-http_gzip_static_module

./configure --prefix=/opt/nginx --user=nginx --group=nginx --with-http_gzip_static_module --with-http_image_filter_module 


#Lua + redis nginx 
./configure --prefix=/opt/nginx --user=nginx --group=nginx --add-module=../ngx_devel_kit-0.2.19 --add-module=../lua-nginx-module-0.9.5rc1 --with-http_gzip_static_module --add-module=../echo --add-module=../lua-resty-redis --with-http_image_filter_module



!!! delete file from html dir --->index.html etc..



