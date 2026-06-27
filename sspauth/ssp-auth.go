package sspauth

import (
	"fmt"
	"strings"

	"github.com/romamalax/libs/convert"
)

const (
	AuthCodeLen = 6
	authMod     = 1_000_000
)

type SSPAuth struct {
	secretA, secretB int
}

func NewSSPAuth(secretA, secretB int) *SSPAuth {
	return &SSPAuth{secretA, secretB}
}

func (a *SSPAuth) Code(sspID int) int {
	if sspID <= 0 {
		return 0
	}
	return (sspID*a.secretA + a.secretB) % authMod
}

func (a *SSPAuth) BuildReqURL(baseURL string, sspID int) string {
	baseURL = strings.TrimRight(baseURL, "/")
	return fmt.Sprintf("%s?sspid=%d&auth=%0*d", baseURL, sspID, AuthCodeLen, a.Code(sspID))
}

func (a *SSPAuth) Validate(sspID int, code []byte) bool {
	if sspID <= 0 || len(code) != AuthCodeLen {
		return false
	}
	return a.Code(sspID) == convert.BytesToInt(code)
}
