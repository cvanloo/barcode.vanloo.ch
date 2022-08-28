server {
	listen 80;
	listen [::]:80;

	server_name barcode.vanloo.ch;

	location / {
		proxy_set_header Host $host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto $scheme;
		proxy_pass http://localhost:8000;
	}

	location /api {
		proxy_set_header Host $host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto $scheme;
		proxy_pass http://localhost:8080/api;
	}

	#location / {
	#	root /var/www/html/barcode.vanloo.ch/;
	#	try_files $uri @backend;
	#}

	#location @backend {
	#	proxy_pass http://127.0.0.1:8000;
	#	proxy_set_header Host $host;
	#	proxy_set_header X-Real-IP $remote_addr;
	#	proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	#	proxy_set_header X-Forwarded-Proto $scheme;
	#}
}
