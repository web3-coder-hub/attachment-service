install:
	go get -u \
	    google.golang.org/protobuf/cmd/protoc-gen-go \
	    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        github.com/go-bindata/go-bindata/go-bindata


clean:
	rm -rf ./third_party/openapiv2/*


generate:
	protoc -I ./proto \
      --go_out ./proto \
      --go_opt paths=source_relative \
      --go-grpc_out ./proto \
      --go-grpc_opt paths=source_relative \
      --grpc-gateway_out ./proto \
      --grpc-gateway_opt paths=source_relative \
      --openapiv2_opt logtostderr=true \
      --openapiv2_opt use_go_templates=true \
      --openapiv2_out ./third_party/openapiv2 \
      ./proto/attachment.proto


swagger:
	go-bindata --nocompress -pkg swagger -o ./third_party/openapiv2/ui/data/swagger/datafile.go ./third_party/swagger-ui/...


build:
	make clean;
	make generate;
	make swagger;