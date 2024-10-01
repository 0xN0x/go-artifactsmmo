package artifactsmmo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewRequest(client *ArtifactsMMO, returnedValue any, method string, url string, body any) *HTTPRequest {
	req, _ := http.NewRequest(method, url, nil)

	return &HTTPRequest{
		Client:        client,
		Body:          body,
		req:           req,
		returnedValue: returnedValue,
	}
}

type HTTPRequest struct {
	Client        *ArtifactsMMO
	Body          any
	req           *http.Request
	returnedValue any
}

func (hc *HTTPRequest) Set() *HTTPRequest {
	// Set headers
	hc.req.Header.Add("Accept", "application/json")
	hc.req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", hc.Client.token))

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

	res, err := hc.Client.client.Do(hc.req)

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
		return nil, fmt.Errorf("failed to parse token: token is missing or empty")
	case 486:
		return nil, fmt.Errorf("an action is already in progress by your character")
	case 497:
		return nil, fmt.Errorf("character inventory is full")
	case 498:
		return nil, fmt.Errorf("character not found")
	case 499:
		return nil, fmt.Errorf("character in cooldown")
	}

	body, _ := io.ReadAll(res.Body)

	json.Unmarshal(body, &hc.returnedValue)

	return res, err
}
