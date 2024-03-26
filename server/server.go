package server

import (
	"attachment-service/cmd"
	"attachment-service/internal/middlewares"
	"attachment-service/internal/utils"
	attachmentPb "attachment-service/proto"
	"attachment-service/server/svc"
	"attachment-service/third_party/openapiv2/ui/data/swagger"
	"context"
	"crypto/tls"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"path"
	"strings"
)

func NewAttachmentService() *svc.AttachmentServer {
	return &svc.AttachmentServer{}
}

func NewServer(conf cmd.ServerConfig, tls *tls.Config) *http.Server {
	grpcServer := newGrpc(conf)
	gwmux, err := newGateway(conf)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(mux)

	return &http.Server{
		Addr:    conf.EndPoint,
		Handler: utils.GrpcHandlerFunc(grpcServer, mux),
		//TLSConfig: tls,
	}
}

func newGrpc(c cmd.ServerConfig) *grpc.Server {
	//creds, err := credentials.NewServerTLSFromFile(c.CertPemPath, c.CertKeyPath)
	//if err != nil {
	//	panic(err)
	//}

	opts := []grpc.ServerOption{
		//grpc.Creds(creds),
		//grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
		grpc_middleware.WithUnaryServerChain(
			middlewares.GrpcRecover(),
			middlewares.GrpcSetContextLogger(),
			middlewares.GrpcAccessLog(),
		),
	}

	server := grpc.NewServer(opts...)

	attachmentPb.RegisterAttachmentServiceServer(server, NewAttachmentService())

	return server
}

func newGateway(c cmd.ServerConfig) (http.Handler, error) {
	ctx := context.Background()
	//_, cancel := context.WithCancel(ctx)
	//defer cancel()

	//dcreds, err := credentials.NewClientTLSFromFile(c.CertPemPath, c.CertServerName)
	//if err != nil {
	//	return nil, err
	//}
	dopts := []grpc.DialOption{
		//grpc.WithTransportCredentials(dcreds),
		grpc.WithInsecure(),
	}

	gwmux := runtime.NewServeMux(
		runtime.WithErrorHandler(runtime.DefaultHTTPErrorHandler),
		runtime.WithStreamErrorHandler(runtime.DefaultStreamErrorHandler),
	)
	if err := attachmentPb.RegisterAttachmentServiceHandlerFromEndpoint(ctx, gwmux, c.EndPoint, dopts); err != nil {
		return nil, err
	}

	return gwmux, nil
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(cmd.ServerConf.SwaggerDir, p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
