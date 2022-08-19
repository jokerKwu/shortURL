package nShortURL

import (
	"fmt"
	"main/aws/ssm"
)

var NShortURLClientID string
var NShortURLSecretKey string

func InitNShortURL() error {
	info, err := ssm.AwsGetParams([]string{"naver_short_url_client_id", "naver_short_url_secret_key"})
	if err != nil {
		return err
	}
	NShortURLClientID = info[0]
	NShortURLSecretKey = info[1]
	fmt.Println("네이버 id,key 초기화 완료")
	fmt.Println(NShortURLClientID)
	fmt.Println(NShortURLSecretKey)
	return nil
}
