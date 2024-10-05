package client

import "net/http"

type ArtifactsConfig struct {
	client   *http.Client
	token    string
	username string
}

func NewConfig(client *http.Client, token, username string) *ArtifactsConfig {
	return &ArtifactsConfig{
		client:   client,
		token:    token,
		username: username,
	}
}

func (ac *ArtifactsConfig) GetClient() *http.Client {
	return ac.client
}

func (ac *ArtifactsConfig) GetToken() string {
	return ac.token
}

func (ac *ArtifactsConfig) GetUsername() string {
	return ac.username
}
