package plugins

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/gogf/gf/v2/util/gconv"

	"time"
)

type AwsCloudfront struct {
	DistId *string   //ID
	Items  []*string //Items
}

func CloudFrontCreateInvalidation(awsCloudfront *AwsCloudfront) error {
	front := cloudfront.New(sess)
	cf := &cloudfront.CreateInvalidationInput{
		DistributionId: awsCloudfront.DistId,
		InvalidationBatch: &cloudfront.InvalidationBatch{
			CallerReference: aws.String(gconv.String(time.Now().Unix())),
			Paths: &cloudfront.Paths{
				Items:    awsCloudfront.Items,
				Quantity: aws.Int64(gconv.Int64(len(awsCloudfront.Items))),
			},
		},
	}
	_, err := front.CreateInvalidation(cf)
	return err
}
