# Simple-echo-server

---
> For studying k8s, pulumi, postgresql...

> It serves http://localhost:80/{string} and will return  {"from_server": "{string}"}
## How to install
### From source
```shell
git clone https://github.com/kvngho/simple-echo-server.git
cd simple-echo-server
docker build -t simple-echo-server .
docker run --rm -p 80:8008 simple-echo-server
```
### From ghcr.io
```shell
docker pull ghcr.io/kvngho/simple-echo-server:latest
docker run --rm -p 80:8008 simple-echo-server
```

## Usage
```shell
curl -X GET "http://localhost:80/helloworld"
```
## Result
```json
{"from_server": "helloworld"}
```