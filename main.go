package refresher

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/oauth2"
)

type Refresher struct {
	/*
		OAuth2 configuration
	*/
	conf *oauth2.Config

	/*
		Refresh token
	*/
	refresh string

	/*
		Access token
	*/
	access string

	/*
		Access token expiration timestamp
	*/
	expire time.Time
}

func New(p, r, c string, e time.Time) (*Refresher, error) {
	var ret Refresher

	if len(p) == 0 {
		return nil, fmt.Errorf("No OAuth2 provider specified")
	}

	switch p {
	case "microsoft":
		ret.conf = &oauth2.Config{
			ClientID:     "08162f7c-0fd2-4200-a84a-f25a4db0b584",
			ClientSecret: "TxRBilcHdC6WGBee]fs?QR:SJ8nI[g82",
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
				TokenURL: "https://login.microsoftonline.com/common/oauth2/v2.0/token",
			},
		}
	case "google":
		return nil, fmt.Errorf("OAuth2 provider Google not yet implemented")
	default:
		return nil, fmt.Errorf("OAuth2 provider %s not supported", p)
	}

	if len(r) == 0 {
		return nil, fmt.Errorf("No refresh token specified")
	}
	ret.refresh = r

	if len(c) > 0 {
		ret.access = c
		ret.expire = e
	} else {
		ret.expire = time.Now().Add(-1 * time.Minute)
	}

	return &ret, nil
}

func (r *Refresher) GetToken() (ret string, err error) {
	if r.expire.Before(time.Now().Add(10 * time.Second)) {
		err = r.update()
		if err != nil {
			return
		}
	}
	ret = r.access
	return
}

func (r *Refresher) GetExpire() time.Time {
	return r.expire
}

func (r *Refresher) update() error {
	oldToken := &oauth2.Token{
		AccessToken:  r.access,
		Expiry:       r.expire,
		RefreshToken: r.refresh,
	}

	newToken, err := r.conf.TokenSource(context.Background(), oldToken).Token()
	if err != nil {
		return err
	}

	if newToken.AccessToken != r.access {
		r.access = newToken.AccessToken
		r.expire = newToken.Expiry
	}

	return nil
}
