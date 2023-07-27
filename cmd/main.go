package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
	"github.com/rs/zerolog/log"
)

func main() {
	verbose := flag.Bool("v", true, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8080", "proxy listen address")
	realm := os.Getenv("REALM_KEY")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	go httpServer()
	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose
	auth.ProxyBasic(proxy, realm, func(user, passwd string) bool {
		return user == username && passwd == password
	})

	log.Info().Str("addr", *addr).Bool("verbose", *verbose).Msg("listen and serve proxy server")
	err := http.ListenAndServe(*addr, proxy)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot listen proxy server")
	}
}

func httpServer() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OKE"))
	})

	log.Info().Str("addr", "8081").Msg("listen and serve http server")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal().Err(err).Msg("cannot listen http server")
	}

}
