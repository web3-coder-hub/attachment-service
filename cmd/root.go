package cmd

import (
	config "attachment-service/internal/utils"
	"flag"
)

type ServerConfig struct {
	CertServerName string
	CertPemPath    string
	CertKeyPath    string
	SwaggerDir     string
	EndPoint       string
	ServerPort     int
}

var ServerConf = new(ServerConfig)

func Execute() {
	flag.IntVar(&ServerConf.ServerPort, "port", config.Conf.Port, "server port")
	flag.StringVar(&ServerConf.CertPemPath, "cert-pem", "./internal/certs/server.pem", "cert-pem path")
	flag.StringVar(&ServerConf.CertKeyPath, "cert-key", "./internal/certs/server.key", "cert-key path")
	flag.StringVar(&ServerConf.SwaggerDir, "swagger-dir", "third_party/openapiv2", "swagger-dir path")
	flag.StringVar(&ServerConf.CertServerName, "cert-server-name", "127.0.0.1", "cert server's hostname")
	flag.Parse()
}
