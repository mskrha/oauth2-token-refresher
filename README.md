[![Go Report Card](https://goreportcard.com/badge/github.com/mskrha/oauth2-token-refresher)](https://goreportcard.com/report/github.com/mskrha/oauth2-token-refresher)

## oauth2-token-refresher

### Description
Simple Go library used to get the OAuth2 access token from the refresh token.

### Installation
`go get github.com/mskrha/oauth2-token-refresher`

### Usage
```go
package main

import (
	"fmt"
	"time"

	"github.com/mskrha/oauth2-token-refresher"
)

func main() {
	o2r, err := refresher.New("microsoft", "Stored refresh token", "Optionally cached access token", time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}
	t, err := o2r.GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Token: %v\n", t)
}
```
