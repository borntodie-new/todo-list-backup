package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailCodeService_SendCode(t *testing.T) {
	var (
		to      = "jasonborntodie@gmail.com"
		code    = "123456"
		subject = "test subject"
	)
	emailService := NewEmailCodeService(to, code, subject)
	err := emailService.SendCode()
	assert.Nil(t, err)
}
