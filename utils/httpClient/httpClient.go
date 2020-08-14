package httpClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type ResponseData struct {
	Data json.RawMessage `json:"data,omitempty"`
}
type ResponseStatus struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_msg,omitempty"`
}

// Response API response struct
type Response struct {
	ResponseStatus
	ResponseData
}

type ResponseMap map[string]interface{}

type HttpClient interface {
	Get(url string, data interface{}) error
	Post(url string, body interface{}, data interface{}) error
	HandleAPIError(res interface{}) error
	Decode(body io.ReadCloser, res interface{}) error
	Unmarshal(byte interface{}, data interface{}) error
}

type httpClient struct{}

func New() *httpClient {
	return &httpClient{}
}

//Get send GET HTTP request
func (h *httpClient) Get(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "http.Get(url)")
	}
	defer resp.Body.Close()

	var res Response
	if err = h.Decode(resp.Body, &res); err != nil {
		return err
	}
	if err = h.HandleAPIError(res); err != nil {
		return err
	}
	if err = h.Unmarshal(res, data); err != nil {
		return err
	}
	return nil
}

//Post send POST HTTP request
func (h *httpClient) Post(url string, body interface{}, data interface{}) error {
	var res ResponseMap
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return errors.Wrapf(err, "json.Marshal(body). body: %+v", body)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return errors.Wrapf(err, "http.Post(url, \"application/json\", bytes.NewBuffer(jsonValue)). jsonValue: %+v", jsonValue)
	}
	defer resp.Body.Close()
	if err = h.Decode(resp.Body, &res); err != nil {
		return err
	}
	if err = h.HandleAPIError(res); err != nil {
		return err
	}
	if err = h.Unmarshal(res, data); err != nil {
		return err
	}
	return nil
}

func (h *httpClient) Decode(body io.ReadCloser, res interface{}) error {
	err := json.NewDecoder(body).Decode(res)
	if err != nil {
		return errors.Wrapf(err, "json.NewDecoder(resp.Body).Decode(&res)")
	}
	return nil
}

func (h *httpClient) Unmarshal(res interface{}, data interface{}) error {
	switch v := res.(type) {
	case Response:
		err := json.Unmarshal(v.Data, &data)
		if err != nil {
			return errors.Wrapf(err, "json.Unmarshal. data: %s", string(v.Data))
		}
	case ResponseMap:
		var r interface{}
		if v["data"] != nil {
			r = v["data"]
		}
		r = v
		bs, err := json.Marshal(r)
		if err != nil {
			return errors.Wrapf(err, "json.Marshal(r). v: %+v", r)
		}
		err = json.Unmarshal(bs, &data)
		if err != nil {
			return errors.Wrapf(err, "json.Unmarshal. data: %s", string(bs))
		}
	}

	return nil
}

func (h *httpClient) HandleAPIError(res interface{}) error {
	switch v := res.(type) {
	case Response:
		if !v.Success {
			return fmt.Errorf("API error: %s", v.ErrorMessage)
		}
	case ResponseMap:
		if v["success"] == false {
			return fmt.Errorf("API error: %s", v["error_msg"])
		}
	}
	return nil
}
