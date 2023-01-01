package nShortURL

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetShortURL(apiURL string) (ResShortURL, error) {
	url := "https://openapi.naver.com/v1/util/shorturl?url=" + apiURL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResShortURL{}, err
	}

	//필요시 헤더 추가 가능
	req.Header.Add("X-Naver-Client-Id", NShortURLClientID)
	req.Header.Add("X-Naver-Client-Secret", NShortURLSecretKey)

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ResShortURL{}, err
	}
	defer resp.Body.Close()

	// 결과 출력
	bytes, _ := ioutil.ReadAll(resp.Body)

	var shortURLObj = ResShortURL{}
	err = json.Unmarshal(bytes, &shortURLObj)
	fmt.Println(shortURLObj)
	if err != nil {
		return ResShortURL{}, err
	}
	if shortURLObj.Code != "200" {
		return ResShortURL{}, fmt.Errorf("%s", shortURLObj.Message)
	}
	if shortURLObj.Result.OrgUrl != apiURL {
		return ResShortURL{}, fmt.Errorf("request url and original url do not match.")
	}
	return shortURLObj, nil
}
