# Athenz N-Token Daemon library [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://opensource.org/licenses/Apache-2.0) [![release](https://img.shields.io/github/release/kpango/ntokend.svg?style=flat-square)](https://github.com/kpango/ntokend/releases/latest) [![CircleCI](https://circleci.com/gh/kpango/ntokend.svg)](https://circleci.com/gh/kpango/ntokend) [![codecov](https://codecov.io/gh/kpango/ntokend/branch/master/graph/badge.svg)](https://codecov.io/gh/kpango/ntokend) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/0045479d5b00466cba2c6650bea2b75f)](https://www.codacy.com/app/i.can.feel.gravity/ntokend?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=kpango/ntokend&amp;utm_campaign=Badge_Grade) [![Go Report Card](https://goreportcard.com/badge/github.com/kpango/ntokend)](https://goreportcard.com/report/github.com/kpango/ntokend) [![GolangCI](https://golangci.com/badges/github.com/kpango/ntokend.svg?style=flat-square)](https://golangci.com/r/github.com/kpango/ntokend) [![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/kpango/ntokend) [![GoDoc](http://godoc.org/github.com/kpango/ntokend?status.svg)](http://godoc.org/github.com/kpango/ntokend)

A daemon that generate and cache Athenz n-token in background.

## Usage

```go
package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"
)
import ntokend "github.com/kpango/ntokend"

func main() {
	keyData, err := ioutil.ReadFile("./private_key.pem")
	if err != nil && keyData == nil {
		log.Fatal(err)
	}

	ntok, err := ntokend.New(
		// load ntoken from file
		ntokend.TokenFilePath(""),
		// validate the ntoken before return
		ntokend.EnableValidate(),
		ntokend.DisableValidate(),
		// ntoken expiry
		ntokend.TokenExpiration(30*time.Minute),
		// ntoken refersh period
		ntokend.RefreshDuration(25*time.Minute),
		// Athenz
		ntokend.AthenzDomain("domain"),
		ntokend.ServiceName("service"),
		ntokend.KeyVersion("keyID"),
		ntokend.KeyData(keyData),
		// ntokend.Hostname("localhost"),
		// ntokend.IPAddr("127.0.0.1"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// start deamon
	ntok.StartTokenUpdater(context.Background())

	// check token exist
	for !ntok.TokenExists() {
		log.Printf("ntoken.TokenExists: false, wait 100ms...")
		time.Sleep(100 * time.Millisecond)
	}
	// get ntoken
	ntoken, err := ntok.GetTokenProvider()()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ntoken: %s", ntoken)

	// force ntoken re-generation
	err = ntok.Update()
}
```
