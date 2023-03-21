package utils

import (
	"crypto/sha512"
	"fmt"
	"hash"
	"strings"
	"sync"

	"github.com/anaskhan96/go-password-encoder"
)

var (
	pwdOnce sync.Once
	pwd     *Crypto
)

type Crypto struct {
	Options *password.Options
}

func NewPWD(saltLen, iterations, keyLen int, fn func() hash.Hash) *Crypto {
	options := &password.Options{
		SaltLen:      saltLen,    // 盐值长度
		Iterations:   iterations, //
		KeyLen:       keyLen,     // 生成的密码长度
		HashFunction: fn,         // 加密算法
	}
	return &Crypto{Options: options}
}

func Default() *Crypto {
	pwdOnce.Do(func() {
		options := &password.Options{
			SaltLen:      8,   // 盐值长度
			Iterations:   100, //
			KeyLen:       32,  // 生成的密码长度
			HashFunction: sha512.New,
		}
		pwd = &Crypto{Options: options}
	})
	return pwd
}

func (p *Crypto) GenPassword(rawPwd string) string {
	salt, encodedPwd := password.Encode(rawPwd, p.Options)
	return fmt.Sprintf("%s$%s", salt, encodedPwd)
}

func (p *Crypto) Verify(rawPwd, encodedPwd string) bool {
	temp := strings.Split(encodedPwd, "$")
	if len(temp) != 2 {
		return false
	}
	return password.Verify(rawPwd, temp[0], temp[1], p.Options)
}
