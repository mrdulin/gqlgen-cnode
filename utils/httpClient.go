package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type ResponseData struct {
	Data interface{} `json:"data"`
}
type ResponseStatus struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_msg"`
}

// Response cnode API response struct
type Response struct {
	ResponseStatus
	ResponseData
}

type ResponseMap map[string]interface{}

type IHttpClient interface {
	Get(url string) (interface{}, error)
	Post(url string, body interface{}) (interface{}, error)
	HandleAPIError(res ResponseMap) error
}

type httpClient struct {
	IHttpClient
}

func NewHttpClient() *httpClient {
	return &httpClient{}
}

// RequestGet send GET HTTP request
func (h *httpClient) Get(url string) (interface{}, error) {
	var res ResponseMap
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "http.Get(url)")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "ioutil.ReadAll(resp.Body). resp.Body: %+v", resp.Body)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, errors.Wrapf(err, "json.Unmarshal error. body: %s", string(body))
	}
	if err := h.HandleAPIError(res); err != nil {
		return nil, err
	}
	return res["data"], nil
}

// RequestPost send POST HTTP request
func (h *httpClient) Post(url string, body interface{}) (interface{}, error) {
	var res ResponseMap
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrapf(err, "json.Marshal(body). body: %+v", body)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, errors.Wrapf(err, "http.Post(url, \"application/json\", bytes.NewBuffer(jsonValue)). jsonValue: %+v", jsonValue)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "ioutil.ReadAll(resp.Body). resp.Body: %+v", resp.Body)
	}

	err = json.Unmarshal(respBody, &res)
	if err != nil {
		return nil, errors.Wrapf(err, "json.Unmarshal. body: %s", string(respBody))
	}
	if err := h.HandleAPIError(res); err != nil {
		return nil, err
	}
	return res, nil
}

func (h *httpClient) HandleAPIError(res ResponseMap) error {
	if res["success"] == false {
		return fmt.Errorf("API error: %+v", res["error_msg"])
	}
	return nil
}
