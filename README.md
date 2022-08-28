# Online Barcode Generator

## Setup

```sh
cp barcode.vanloo.ch /etc/nginx/sites-available/
ln -s /etc/nginx/sites-available/barcode.vanloo.ch /etc/nginx/sites-enabled/

docker-compose up -d --no-deps --build
```

### SELinux policies

```sh
sudo grep nginx /var/log/audit/audit.log | grep denied | audit2allow -M nginxlocalconf
sudo semodule -i nginxlocalconf.pp 
```
