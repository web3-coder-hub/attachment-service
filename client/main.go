package main

import (
	"log"

	attachmentPb "attachment-service/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	//creds, err := credentials.NewClientTLSFromFile("./internal/certs/server.pem", "localhost")
	//if err != nil {
	//	log.Printf("Failed to create TLS credentials %v", err)
	//}
	conn, err := grpc.Dial(
		":9001",
		//grpc.WithTransportCredentials(creds),
		grpc.WithInsecure(),
	)
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := attachmentPb.NewAttachmentServiceClient(conn)
	context := context.Background()
	body := &attachmentPb.PDFRequest{}

	r, err := c.CreatePDF(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Data.OutFileUrl)
}
