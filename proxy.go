package main

import (
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
)


func main() {
	staticProxyToken, exists := os.LookupEnv("STATIC_PROXY_TOKEN")
    if !exists {
    	panic("STATIC_PROXY_TOKEN is required")
    }
	authorize := func(user, passwd string) bool {
		return passwd == staticProxyToken
	}
	proxy := goproxy.NewProxyHttpServer()
	auth.ProxyBasic(proxy, "proxy server", authorize)
    log.Fatal(http.ListenAndServe(":8081", proxy))
}






