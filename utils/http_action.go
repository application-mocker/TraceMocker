package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func DoBlockRequestWithJson(method string, url string, param, header map[string]string, body interface{}) (*http.Response, error) {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyBytesBuffer := bytes.NewReader(bodyJson)
	if param == nil {
		param = map[string]string{}
	}
	if header == nil {
		param = map[string]string{}
	}

	return DoBlockRequest(method, url, param, header, bodyBytesBuffer)
}

func DoBlockRequest(method string, url string, param, header map[string]string, body io.Reader) (*http.Response, error) {
	req, err := NewRequest(method, url, param, header, body)
	if err != nil {
		return nil, err
	}
	return (&http.Client{}).Do(req)
}

func NewRequest(method string, url string, param, header map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	for key, value := range param {
		if query.Has(key) {
			query.Del(key)
		}
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	for key, value := range header {
		query.Set(key, value)
	}
	return req, nil
}
