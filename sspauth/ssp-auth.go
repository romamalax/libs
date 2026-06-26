package sspauth

import "github.com/romamalax/libs/convert"

const lenAuthCode = 6

type SSPAuth struct {
	secretA, secretB int
}

func NewSSPAuth(secretA, secretB int) *SSPAuth {
	return &SSPAuth{secretA, secretB}
}

func (a *SSPAuth) Validate(sspID int, code []byte) bool {
	if len(code) != lenAuthCode {
		return false
	}
	if (sspID*a.secretA+a.secretB)%1_000_000 == convert.BytesToInt(code) {
		return true
	}
	return false
}
