package common

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"bitbucket.org/applysquare/applysquare-go/pkg/base/trace"
)

func HttpGetBytes(c *http.Client, url string) ([]byte, error) {
	// glog.Infof("GET %s", url)
	resp, err := c.Get(url)
	if resp != nil {
		defer resp.Body.Close()
		defer io.Copy(ioutil.Discard, resp.Body) // Ensure body is exhausted, so connection can be reused.
	}
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Http error: %d, %s", resp.StatusCode, resp.Status)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func Post(tr trace.T, urlString string, params map[string]string, headers map[string]string, response interface{}) error {
	tr.Printf("post url: %s", urlString)

	reqBody := url.Values{}
	for key, value := range params {
		reqBody.Add(key, value)
	}
	req, err := http.NewRequest("POST", urlString, strings.NewReader(reqBody.Encode()))
	if err != nil {
		tr.Printf("err: %v", err)
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		tr.Printf("err: %v", err)
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		tr.Printf("err: %v", err)
		return err
	}

	if err = json.Unmarshal(respBody, &response); err != nil {
		tr.Printf("err: %v", err)
		return err
	}
	tr.Printf("done")
	return nil
}
