package service

import (
	"context"

	"golang.org/x/oauth2"
)

func GetAuthLink(config *oauth2.Config) string {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return authURL
}

func GetToken(config *oauth2.Config, authCode string) (*oauth2.Token, error) {
	token, err := config.Exchange(context.TODO(), authCode)

	if err != nil {
		return nil, err
	}

	return token, nil
}
