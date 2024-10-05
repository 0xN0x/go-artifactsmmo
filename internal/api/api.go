package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/0xN0x/go-artifactsmmo/internal/client"
	"github.com/0xN0x/go-artifactsmmo/models"
)

func NewRequest(config *client.ArtifactsConfig, customStruct any, method string, url string, body any) *HTTPRequest {
	req, _ := http.NewRequest(method, url, nil)

	return &HTTPRequest{
		Config:       config,
		Body:         body,
		req:          req,
		customStruct: customStruct,
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

func (hc *HTTPRequest) AddBody() *HTTPRequest {
	if hc.Body != nil {
		// Set body & marshal
		marshalled, err := json.Marshal(hc.Body)
		if err != nil {
			panic(err)
		}

		hc.req.Body = io.NopCloser(bytes.NewBuffer(marshalled))
		hc.req.ContentLength = int64(len(marshalled))
	}

	return hc
}

func (hc *HTTPRequest) Run() (*http.Response, error) {
	hc.Set().AddBody()

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
	case 478:
		return nil, models.ErrInsufficientQuantity
	case 486:
		return nil, models.ErrActionInProgress
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
