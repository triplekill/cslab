package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256HMAC(payload, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(payload))
	return h.Sum(nil)
}

func RandomString(length int) string {
	buf := make([]byte, length/2+1)
	rand.Read(buf)
	ret := hex.EncodeToString(buf)
	return ret[:length]
}
