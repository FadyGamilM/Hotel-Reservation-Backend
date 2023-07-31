package main

import "flag"

type AppConfig struct {
	ListenAddr string
}

func NewAppConfig() *AppConfig {
	listenAddr := flag.String("listenAddr", ":5000", "the port of our server api")
	flag.Parse()
	config := AppConfig{ListenAddr: *listenAddr}
	return &config
}
