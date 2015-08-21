CDN Config for Nginx

Gereksinimler

Centon 6.6+ (x64)

daemonize Tool
yum install -y http://puias.math.ias.edu/data/puias/unsupported/6/x86_64/daemonize-1.7.5-6.sdl6.x86_64.rpm

Kurulum

cp /etc/rc.d/init.d/mstrcdn
cp /bin/mstrcdn
mkdir /etc/maestropanel/agent/config
mkdir /etc/maestropanel/nginx/tmpl
cp /etc/maestropanel/agent/config/mstrcdn.conf
cp /etc/maestropanel/nginx/tmpl/*


Konfigürasyon

/etc/maestropanel/agent/config/mstrcdn.conf


API


POST /Cdn/Create

Yeni Nginx konfigürasyonu oluşturur.



Parametreler

name: domain name
ipaddr: vhosts'un çalışacağı IP adresi
port: vhost'un portu (default 80)
ssl: Vhost'u SSL özelliği içeren template'i ile açar.
full: Vhost'u Split olmayan template ile açar.

İstek:

http://localhost:9721/Cdn/Create


Cevap:

{
    "success": true,
    "message": "Build Success: D:\\nginx\\vhosts\\osman.com.conf"
}


DELETE /Cdn/Delete



GET /Cdn/List


Cevap:

{
    "success": true,
    "message": "Success",
    "vhosts": [
        "D:\\nginx\\vhosts",
        "D:\\nginx\\vhosts\\domain.com.conf"
		"D:\\nginx\\vhosts\\osman.com.conf"
    ]
}