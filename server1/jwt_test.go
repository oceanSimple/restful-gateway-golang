package main

import (
	"server1/tool"
	"testing"
)

func TestGetJwt(t *testing.T) {
	token, err := tool.GetJwtToken(tool.JwtPayLoad{}, "ocean1234567890", 3600)
	if err != nil {
		t.Error("Failed to get jwt token")
	} else {
		t.Log("Token:", token)
	}
}
