server {
    # NOTE:
    # You might need to add the following lines to /etc/nginx/nginx.conf:
    # include /etc/nginx/mime.types;
    # types {
    #    application/wasm wasm;
    # }
    # default_type application/octet-stream;

    server_name barcode.vanloo.ch;
    root /var/www/html/barcode.vanloo.ch;
    index index.html;

    # Let Certbot setup SSL
	listen 80;
	listen [::]:80;
}
