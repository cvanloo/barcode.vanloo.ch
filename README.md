# Online Barcode Generator

A simple website to generate barcodes.

- Create many barcodes and list them side by side
- Barcodes are generated on the client side using WASM
- Sessions are persisted to local storage
- Restore previous sessions
- Drag'n'Drop or delete barcodes

## Pre Commit

```sh
pre-commit install
```

NOTE: Changes to `.pre-commit-config.yaml` require `--no-verify` in order
to be committed:

```sh
git add .pre-commit-config.yaml
git commit --no-verify
```

## Setup

```sh
cp barcode.vanloo.ch /etc/nginx/sites-available/
ln -s /etc/nginx/sites-available/barcode.vanloo.ch /etc/nginx/sites-enabled/

cd api/
make install

sudo cp web/public/* /var/www/html/barcode.vanloo.ch/
sudo chown www:www /var/www/html/barcode.vanloo.ch/*
```

### SELinux policies

```sh
sudo grep nginx /var/log/audit/audit.log | grep denied | audit2allow -M nginxlocalconf
sudo semodule -i nginxlocalconf.pp 
```
