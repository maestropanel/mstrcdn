MaestroPanel CDN Agent
======================
Nginx için http RestAPI üzerinden otomatik konfigürasyon oluşturan servistir.

> Desteklenen Dağıtım Centon 6.6 ve üzeri (x64)

Gereksinimler
-------------

**daemonize**

    yum install -y http://puias.math.ias.edu/data/puias/unsupported/6/x86_64/daemonize-1.7.5-6.sdl6.x86_64.rpm

Kurulum
-------
Dizin yapısı

/etc/rc.d/init.d/mstrcdn
init dosyasıdır. Servisi sisteme register eder.

/bin/mstrcdn 
Ana dosyadır. Servisin core dosyasıdır.

/etc/maestropanel/agent/config/mstrcdn.conf
Servisin çalışacak değişkenlerinin tutulduğu konfigürasyon dosyasıdır. Standart ini formatındadır. 

/etc/maestropanel/nginx/tmpl 
Nginx konfigürasyon templatelerinin tutulduğu klasödür. Servis 3 isimde template dosyası tanımaktadır. (full.cdn.tmpl, split.cdn.tmpl, ssl.cdn.tmpl)

Konfigürasyon
-------------
Servisin çalışması için gerekli olan konfidürasyon dosyasına aşağıdaki dizinden ulaşabilirsiniz.

    /etc/maestropanel/agent/config/mstrcdn.conf

*mstrcdn.conf*
    [api]
    Port=9722
    SecretKey=qK624M3ZrpfCrlia5jQn
    ConfigRoot="/usr/local/nginx/conf/sites"
    TemplatePath="/etc/maestropanel/nginx/tmpl"

**Port**
Servisin hangi portu dinlyeceğini belirler.

**SecretKey**
Servise ulaşılırken Authentication header'ında kullanılacak gizli anahtarı belirler.

**ConfigRoot**
Nginx için oluşturulacak konfigürasyon dosyalarının yolunu belirler.

**TemplatePath**
Nginx için önceden belirlenmeiş şablon konfigürasyon dosyalarının yolunu belirler.

----------
API
===

Servis HTTP üzerinden Rest olarak çalışmaktadır ve geri dönüş modeli JSON'dur

Geri dönüş modeli;

    { "success": true, "message": "İşlem ile ilgili mesaj" }


POST /Cdn/Create
----------------
Yeni Nginx konfigürasyonu oluşturur.

**Parametreler**

 - name: domain name  
 - ipaddr: vhosts'un çalışacağı IP adresi  
 - port: vhost'un portu (default 80)  
 - ssl: Vhost'u SSL özelliği içeren template'i ile açar.  
 - full: Vhost'u Split olmayan template ile açar.

DELETE /Cdn/Delete
------------------
Mevcut bir domain'in konfigürasyonu siler.

**Parametreler**

 - name: domain name

GET /Cdn/List
-------------
Mevcut domainlerin listerini verir.

Cevap:

    {
        "success": true,
        "message": "List success",
        "vhosts": [
            "/usr/local/nginx/conf/sites",
            "/usr/local/nginx/conf/sites/cdn.domain.com.tr.conf",
            "/usr/local/nginx/conf/sites/cdn.demo.com.tr.conf",
            "/usr/local/nginx/conf/sites/ssl.cdn.hoppa.com.tr.conf"
        ]
    }
