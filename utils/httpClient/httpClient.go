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
	Data json.RawMessage `json:"data"`
}
type ResponseStatus struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_msg"`
}

// Response API response struct
type Response struct {
	ResponseStatus
	ResponseData
}

type HttpClient interface {
	Get(url string, data interface{}) error
	Post(url string, body interface{}, data interface{}) error
	HandleAPIError(res Response) error
	Decode(body io.ReadCloser, res *Response) error
	Unmarshal(byte []byte, data interface{}) error
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
	if err = h.Unmarshal(res.Data, data); err != nil {
		return err
	}
	return nil
}

//Post send POST HTTP request
func (h *httpClient) Post(url string, body interface{}, data interface{}) error {
	var res Response
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
	if err = h.Unmarshal(res.Data, data); err != nil {
		return err
	}
	return nil
}

func (h *httpClient) Decode(body io.ReadCloser, res *Response) error {
	err := json.NewDecoder(body).Decode(res)
	if err != nil {
		return errors.Wrapf(err, "json.NewDecoder(resp.Body).Decode(&res)")
	}
	return nil
}

func (h *httpClient) Unmarshal(byte []byte, data interface{}) error {
	err := json.Unmarshal(byte, &data)
	if err != nil {
		return errors.Wrapf(err, "json.Unmarshal. []byte: %s", string(byte))
	}
	return nil
}

func (h *httpClient) HandleAPIError(res Response) error {
	if !res.Success {
		return fmt.Errorf("API error: %s", res.ErrorMessage)
	}
	return nil
}
