package lib

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func Md5Password(str string) (string) {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func CreateToken() string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, 32)
	for i := 0; i < 32; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
