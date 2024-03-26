package plugins

import "io"

type Ses struct {
	Recipient string //收件人
	Sender    string //发件人
	HtmlBody  string //HTML文本
	TextBody  string //TEXT文本
	Subject   string //邮件标题
}

type S3 struct {
	BuketName   *string
	KeyName     *string
	ContentType *string
	Body        io.Reader
}

type Cloudfront struct {
	DistId *string   //ID
	Items  []*string //Items
}
