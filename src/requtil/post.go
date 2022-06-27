package requtil

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
	"ton-event-idx/src/logger"
)

func SendPostReq(url string, postData []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Timeout: 2 * time.Second, Transport: tr}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			logger.Errorf("can't close response.Body: %s", err)
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
