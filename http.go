package google

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var UserAgent string = defaultUserAgent

const defaultRequestTimeout = 10 //seconds
var RequestTimeout int = defaultRequestTimeout

func HTTPRequest(ctx context.Context, url string, headers map[string]string) ([]byte, error) {

	client := &http.Client{
		Timeout: time.Second * time.Duration(RequestTimeout),
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d %s", res.StatusCode, res.Status)
	}

	return ioutil.ReadAll(res.Body)
}
