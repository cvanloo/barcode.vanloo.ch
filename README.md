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
docker build -t docker-barcode .
docker run -p 8080:80 -i -t docker-barcode:latest

# test it out
curl localhost:8080
```

### Using Docker Compose

```yaml
services:
    barcode:
        build: https://github.com/cvanloo/barcode.vanloo.ch.git
        container_name: barcode
    caddy:
        image: caddy:2
        container_name: caddy
        ports:
            - "80:80"
            - "443:443"
        volumes:
            - /etc/docker/Caddyfile:/etc/caddy/Caddyfile
            - caddy_data:/data
            - caddy_config:/config
volumes:
    caddy_data:
    caddy_config:
```

...where `/etc/docker/Caddyfile` is something like:

```
barcode.example.com {
    reverse_proxy http://barcode:80
}
```

When setup like this, the reverse proxy handles and terminates TLS.

Run the commands:

```
docker compose -f config.yml build # build it
docker compose -f config.yml up -d # start it up
docker compose -f config.yml down  # tear it down
```

You might also want to setup a systemd service to automatically start when the server boots:

```
[Unit]
Description=Start Docker Services
Requires=docker.service
After=docker.service

[Service]
Restart=always
ExecStart=/usr/bin/docker compose -f /etc/docker/config.yml up
ExecStop=/usr/bin/docker compose -f /etc/docker/config.yml down

[Install]
WantedBy=default.target
```
