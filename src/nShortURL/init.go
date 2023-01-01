package nShortURL

import (
	"main/aws/ssm"
)

var NShortURLClientID string
var NShortURLSecretKey string

type ResShortURL struct {
	Result  ResResult
	Message string
	Code    string
}
type ResResult struct {
	Url    string
	Hash   string
	OrgUrl string
}

func InitNShortURL() error {
	info, err := ssm.AwsGetParams([]string{"naver_short_url_client_id", "naver_short_url_secret_key"})
	if err != nil {
		return err
	}
	NShortURLClientID = info[0]
	NShortURLSecretKey = info[1]
	return nil
}
