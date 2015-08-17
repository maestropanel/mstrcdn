CDN Config for Nginx

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