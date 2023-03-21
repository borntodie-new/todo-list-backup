package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	crypto *Crypto
)

func init() {
	crypto = Default()
}

func TestCrypto_GenPassword(t *testing.T) {
	password := crypto.GenPassword("jason123")
	t.Log(password)
}
func TestCrypto_VerifyWithTrue(t *testing.T) {
	rawPassword := "jason123"
	encodedPassword := crypto.GenPassword(rawPassword)
	verify := crypto.Verify(rawPassword, encodedPassword)
	assert.True(t, verify)
}
func TestCrypto_VerifyWithFalse(t *testing.T) {
	rawPassword := "jason123"
	encodedPassword := crypto.GenPassword(rawPassword)
	verify := crypto.Verify(rawPassword+"1231", encodedPassword)
	assert.False(t, verify)
}
