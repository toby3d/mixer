package oauth

import (
	"net/url"

	oauth2 "golang.org/x/oauth2"
)

type (
	Config struct {
		*oauth2.Config
	}

	Token struct {
		*oauth2.Token
	}
)

func NewClient(clientID, clientSecret, redirectURI string, scopes ...string) *Config {
	auth := &url.URL{
		Scheme: "https",
		Host:   "beam.pro",
		Path:   "/oauth/authorize",
	}

	token := &url.URL{
		Scheme: "https",
		Host:   "beam.pro",
		Path:   "/api/v1/oauth/token",
	}

	return &Config{
		&oauth2.Config{
			RedirectURL:  redirectURI,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  auth.String(),
				TokenURL: token.String(),
			},
		},
	}
}

func (cfg *Config) Exchange(code string) (*Token, error) {
	oauth2.RegisterBrokenAuthHeaderProvider(cfg.Config.Endpoint.TokenURL)
	token, err := cfg.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}
	return &Token{token}, nil
}
