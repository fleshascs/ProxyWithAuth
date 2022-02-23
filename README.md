# Proxy server with a basic Auth

## Run localy

```shell script
STATIC_PROXY_TOKEN=secret_token go run proxy.go

curl -x "http://some-configuration-for-the-future:secret_token@127.0.0.1:8081" "https://www.example.com/"
```

## Run with docker

```shell script
docker build -t fleshas/proxy-with-auth .
docker run -d -e STATIC_PROXY_TOKEN=secret_token -p 8081:8081 --name ms fleshas/proxy-with-auth:latest
```
