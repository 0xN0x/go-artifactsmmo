package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/0xN0x/go-artifactsmmo/internal/client"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func NewRequest(config *client.ArtifactsConfig) *HTTPRequest {
	req, _ := http.NewRequest("", "", nil)

	return &HTTPRequest{
		Config: config,
		req:    req,
	}
}

type HTTPRequest struct {
	Config       *client.ArtifactsConfig
	Body         any
	req          *http.Request
	customStruct any
}

func (hc *HTTPRequest) Set() *HTTPRequest {
	// Set headers
	hc.req.Header.Add("Accept", "application/json")
	hc.req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", hc.Config.GetToken()))

	if hc.req.Method == "POST" {
		hc.req.Header.Add("Content-Type", "application/json")
	}

	return hc
}

func (hc *HTTPRequest) SetResultStruct(customStruct any) *HTTPRequest {
	hc.customStruct = customStruct
	return hc
}

func (hc *HTTPRequest) SetMethod(method string) *HTTPRequest {
	hc.req.Method = method
	return hc
}

func (hc *HTTPRequest) SetURL(uri string) *HTTPRequest {
	u, err := url.Parse(fmt.Sprintf("%s%s", hc.Config.GetApiUrl(), uri))
	if err != nil {
		panic(err)
	}

	hc.req.URL = u
	return hc
}

func (hc *HTTPRequest) SetBody(body any) *HTTPRequest {
	hc.Body = body

	marshalled, err := json.Marshal(hc.Body)
	if err != nil {
		panic(err)
	}

	hc.req.Body = io.NopCloser(bytes.NewBuffer(marshalled))
	hc.req.ContentLength = int64(len(marshalled))

	return hc
}

func (hc *HTTPRequest) SetParam(key string, value string) *HTTPRequest {
	q := hc.req.URL.Query()
	q.Add(key, value)
	hc.req.URL.RawQuery = q.Encode()

	return hc
}

func (hc *HTTPRequest) Run() (*http.Response, error) {
	hc.Set()

	res, err := hc.Config.GetClient().Do(hc.req)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	switch res.StatusCode {
	case 452:
		return nil, models.ErrBadToken
	case 474:
		return nil, models.ErrTaskNotOwned
	case 475:
		return nil, models.ErrTaskAlreadyCompleted
	case 478:
		return nil, models.ErrInsufficientQuantity
	case 486:
		return nil, models.ErrActionInProgress
	case 487:
		return nil, models.ErrCharacterHasNoTask
	case 488:
		return nil, models.ErrTaskNotCompleted
	case 489:
		return nil, models.ErrCharacterAlreadyHasTask
	case 497:
		return nil, models.ErrCharacterFullInventory
	case 498:
		return nil, models.ErrCharacterNotFound
	case 499:
		return nil, models.ErrCharacterInCooldown
	}

	body, _ := io.ReadAll(res.Body)

	var returned = struct {
		Data any `json:"data"`
	}{
		Data: hc.customStruct,
	}

	json.Unmarshal(body, &returned)

	return res, err
}
