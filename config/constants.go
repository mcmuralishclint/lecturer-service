package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "74216694885-buhikcp8866cfc0niudeskt1hjik0m38.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:3000/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
	return conf
}
