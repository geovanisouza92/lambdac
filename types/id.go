package types

import (
	"math/rand"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

func NewID() string {
	return randomString(40)
}

func randomString(n int) string {
	b := make([]byte, n)
	i, cache, remain := n-1, src.Int63(), letterIdxMax
	for i >= 0 {
		if remain == 0 {
			cache = src.Int63()
			remain = letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
