# docker build -t barcode-web -f Dockerfile.web .
# docker run -it --rm --name barcode-web-running -p 8000:8000 barcode-web:latest
FROM caddy:2.5.2-alpine AS webapp
COPY Caddyfile /etc/caddy/Caddyfile
COPY public /srv
EXPOSE 8000
