package plugins

import (
	config "attachment-service/internal/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

func InitAWS(awsConfig *config.AwsConfig) (err error) {
	sess, err = session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(awsConfig.AccessKeyID, awsConfig.SecretAccessKey, ""),
		Region:      aws.String(awsConfig.Region)},
	)
	return
}
