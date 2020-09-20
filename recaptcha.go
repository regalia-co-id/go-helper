package helper

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func CaptchaVerify(urlApi string, secret string, token string) (bool, error) {
	type CaptchaVerifyResponse struct {
		Success bool
	}

	var captchaVerifyResponse CaptchaVerifyResponse
	captchaVerifyResponse.Success = false

	captchaPayloadRequest := url.Values{}
	captchaPayloadRequest.Set("secret", secret)
	captchaPayloadRequest.Set("response", token)

	verifyCaptchaRequest, err := http.NewRequest("POST", urlApi, strings.NewReader(captchaPayloadRequest.Encode()))
	verifyCaptchaRequest.Header.Add("content-type", "application/x-www-form-urlencoded")
	verifyCaptchaRequest.Header.Add("cache-control", "no-cache")
	if err != nil {
		return false, err
	}

	verifyCaptchaResponse, err := http.DefaultClient.Do(verifyCaptchaRequest)
	if err != nil {
		return false, err
	}

	decoder := json.NewDecoder(verifyCaptchaResponse.Body)
	decoderErr := decoder.Decode(&captchaVerifyResponse)

	defer verifyCaptchaResponse.Body.Close()

	if decoderErr != nil {
		return false, decoderErr
	}

	return captchaVerifyResponse.Success, nil
}
