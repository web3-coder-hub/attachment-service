package main

import (
	"attachment-service/cmd"
	logger "attachment-service/internal/middlewares"
	"attachment-service/internal/plugins"
	"attachment-service/internal/utils"
	config "attachment-service/internal/utils"
	"attachment-service/server"
	"fmt"
	"log"
	"net"
)

func init() {
	config.Init("./etc/attachment-api.yaml")
	logger.InitLogger()
	plugins.InitAWS(config.Conf.AwsConfig)
	//grpcLog := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	//grpclog.SetLoggerV2(grpcLog)
}

func main() {
	cmd.Execute()
	conf := cmd.ServerConf

	EndPoint := fmt.Sprintf(":%v", conf.ServerPort)
	tlsConfig := utils.GetTLSConfig(conf.CertPemPath, conf.CertKeyPath)
	conf.EndPoint = EndPoint

	conn, err := net.Listen("tcp", EndPoint)
	if err != nil {
		log.Printf("TCP Listen err:%v\n", err)
	}

	srv := server.NewServer(*conf, tlsConfig)
	log.Printf("Serving gRPC and https on 127.0.0.1:%v\n", conf.ServerPort)

	if err = srv.Serve(conn); err != nil {
		log.Printf("ListenAndServe err: %v\n", err)
	}
	//if err = srv.Serve(utils.NewTLSListener(conn, tlsConfig)); err != nil {
	//	log.Printf("ListenAndServe err: %v\n", err)
	//}
}
