package plugins

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

const CharSet = "UTF-8"

type AwsSes struct {
	Recipient string //收件人
	Sender    string //发件人
	HtmlBody  string //HTML文本
	TextBody  string //TEXT文本
	Subject   string //邮件标题
}

func SaveEmail(awsSes *AwsSes) error {
	// Create an SES session.
	svc := ses.New(sess)
	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				&awsSes.Recipient,
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    &awsSes.HtmlBody,
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    &awsSes.TextBody,
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    &awsSes.Subject,
			},
		},
		Source: &awsSes.Sender,
		// Uncomment to use a configuration set
		//ConfigurationSetName: awsPlugin.String(ConfigurationSet),
	}
	// Attempt to send the email.
	_, err := svc.SendEmail(input)
	return err
}
