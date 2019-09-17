package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/armon/go-socks5"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var config AppConfig

	if err := envconfig.Process("socks", &config); err != nil {
		log.Fatalf("[Config] Critical error: %v", err)
	}
	log.Printf("[Config] Debug: %t", config.Debug)

	conf := &socks5.Config{}

	requireAuth := bool(config.User != "" && config.Password != "")
	log.Printf("[Config] Auth: %t", requireAuth)
	if requireAuth {
		creds := socks5.StaticCredentials{config.User: config.Password}
		cator := socks5.UserPassAuthenticator{Credentials: creds}
		conf.AuthMethods = []socks5.Authenticator{cator}
	}

	log.Printf("[Config] Restict to white list: %t", config.Restrict)
	if config.Restrict {
		var rules Rules
		dat, err := ioutil.ReadFile("rules.json")
		if err != nil {
			log.Fatalf("[Config] Critical error: %v", err)
		}
		if err = json.Unmarshal(dat, &rules); err != nil {
			log.Fatalf("[Config] Critical error: %v", err)
		}
		err = rules.Load()
		if err != nil {
			log.Fatalf("[Config] Critical error: %v", err)
		}
		conf.Rules = &rules
	}

	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf("[Proxy] Critical error: %v", err)
	}

	listenAddr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	socksURL := fmt.Sprintf("socks5://%s:%d", config.Host, config.Port)
	if requireAuth {
		socksURL = fmt.Sprintf("socks5://%s:%s@%s:%d", config.User, config.Password, config.Host, config.Port)
	}

	log.Printf("[Proxy] Listening at %s", listenAddr)
	log.Printf("[Proxy] Use this URL to connect: %s", socksURL)

	if err := server.ListenAndServe("tcp", listenAddr); err != nil {
		log.Fatalf("[Proxy] Critical error: %v", err)
	}
}
