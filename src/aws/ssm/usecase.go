package ssm

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	awsInit "main/aws"
	"strings"
)

func AwsGetParams(paths []string) ([]string, error) {
	ctx := context.TODO()
	// get ssm param
	params, err := awsInit.AwsClientSsm.GetParameters(ctx, &ssm.GetParametersInput{
		Names:          paths,
		WithDecryption: true,
	})
	if err != nil {
		return nil, err
	}
	result := make([]string, len(paths))
	for i, path := range paths {
		val := ""
		for _, parameter := range params.Parameters {
			if strings.Contains(aws.ToString(parameter.ARN), path) {
				val = aws.ToString(parameter.Value)
				break
			}
		}
		result[i] = val
	}
	return result, nil
}

func AwsGetParam(path string) (string, error) {
	ctx := context.TODO()
	// get ssm param
	param, err := awsInit.AwsClientSsm.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(path),
		WithDecryption: true,
	})
	fmt.Println(param.ResultMetadata)
	fmt.Println(param.Parameter.Value)
	if err != nil {
		return "", err
	}
	return aws.ToString(param.Parameter.Value), nil
}
