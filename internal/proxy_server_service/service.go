package proxy_server_service

import (
	"fmt"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
	"github.com/rs/zerolog/log"
)

type proxyServerService struct {
	port     int
	username string
	password string
}

func NewProxyServerService(port int, username string, password string) *proxyServerService {
	return &proxyServerService{
		port:     port,
		username: username,
		password: password,
	}
}

func (service *proxyServerService) Start() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	auth.ProxyBasic(proxy, "whoant", func(user, passwd string) bool {
		return user == service.username && passwd == service.password
	})

	log.Info().Int("port", service.port).Msg("start proxy server ...")
	err := http.ListenAndServe(fmt.Sprintf(":%d", service.port), proxy)
	if err != nil {
		log.Fatal().Err(err).Int("port", service.port).Msg("cannot start proxy server")
	}
}
