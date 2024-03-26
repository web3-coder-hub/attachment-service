package svc

import (
	attachmentPb "attachment-service/proto"
	"fmt"
	"golang.org/x/net/context"
)

func (a AttachmentServer) CreateCSV(ctx context.Context, r *attachmentPb.CSVRequest) (*attachmentPb.CSVResponse, error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Print("panic: ", e)
		}
	}()
	return &attachmentPb.CSVResponse{}, nil
}
