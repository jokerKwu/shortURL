package aws

import (
	"context"
	"fmt"
	AwsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var AwsClientSsm *ssm.Client

func InitAws(region string) error {

	awsConfig, err := AwsConfig.LoadDefaultConfig(context.TODO(),
		AwsConfig.WithRegion(region))
	if err != nil {
		return err
	}
	AwsClientSsm = ssm.NewFromConfig(awsConfig)
	fmt.Println("aws ssm 초기화 완료")
	return nil
}
