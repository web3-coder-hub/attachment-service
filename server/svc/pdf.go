package svc

import (
	Err "attachment-service/internal/commons"
	"attachment-service/internal/plugins"
	config "attachment-service/internal/utils"
	attachmentPb "attachment-service/proto"
	"fmt"
	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/aws/aws-sdk-go/aws"
	"golang.org/x/net/context"
	"os"
	"strings"
)

type AttachmentServer struct {
	attachmentPb.AttachmentServiceServer
}

func (a AttachmentServer) CreatePDF(ctx context.Context, r *attachmentPb.PDFRequest) (*attachmentPb.PDFResponse, error) {
	meta := r.Meta
	html := r.Html
	url := r.Url

	defer func() {
		if e := recover(); e != nil {
			fmt.Print("panic: ", e)
		}
	}()

	if meta == nil || meta.Title == nil || url == "" && html == "" {
		return &attachmentPb.PDFResponse{
			Data:    nil,
			Code:    Err.InvalidParam,
			Message: Err.InvalidParamError.String(),
		}, nil
	}

	pdfg, err := pdf.NewPDFGenerator()
	if err != nil {
		return &attachmentPb.PDFResponse{
			Data:    nil,
			Code:    Err.ServerInternal,
			Message: Err.ServerInternalError.String(),
		}, nil
	}

	pdfg.Orientation.Set(meta.Orientation.Enum().String())

	if meta.Dpi != nil {
		pdfg.Dpi.Set(uint(*meta.Dpi))
	}

	pdfg.Title.Set(*meta.Title)
	pdfg.PageSize.Set(meta.PrintSize.Enum().String())

	if url != "" {
		page := pdf.NewPage(url)
		pdfg.AddPage(page)
	} else {
		pdfg.AddPage(pdf.NewPageReader(strings.NewReader(html)))
	}

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		return &attachmentPb.PDFResponse{
			Data:    nil,
			Code:    Err.ServerInternal,
			Message: Err.ServerInternalError.String(),
		}, nil
	}

	// Write buffer contents to file on disk
	outPath := fmt.Sprintf("./%v.pdf", *meta.Title)
	err = pdfg.WriteFile(outPath)
	if err != nil {
		return &attachmentPb.PDFResponse{
			Data:    nil,
			Code:    Err.ServerInternal,
			Message: Err.ServerInternalError.String(),
		}, nil
	}

	awsS3 := new(plugins.AwsS3)
	fileReader, _ := os.Open(outPath)
	defer os.Remove(outPath)
	defer fileReader.Close()
	s3Config := config.Conf.FileS3Config
	awsS3.BuketName = aws.String(s3Config.BucketName)
	keyName := fmt.Sprintf("/pdf/%v.pdf", *meta.Title)
	awsS3.KeyName = &keyName
	awsS3.Body = fileReader
	out, uploadErr := plugins.UploaderS3PublicRead(awsS3)
	if uploadErr != nil {
		return &attachmentPb.PDFResponse{
			Data:    nil,
			Code:    Err.ServerInternal,
			Message: Err.ServerInternalError.String(),
		}, nil
	}
	return &attachmentPb.PDFResponse{
		Data: &attachmentPb.PDFData{
			OutFileUrl: out.Location,
		},
		Code:    Err.Sucess,
		Message: "",
	}, nil
}
