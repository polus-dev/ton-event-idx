package spreq

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

func SendJsonPostReq(url string, timeout time.Duration, postData []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Timeout: timeout, Transport: tr}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func() { response.Body.Close() }()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
