package hashext

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Md5Str(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	return hex.EncodeToString(bs)
}

func Sha256Str(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	return hex.EncodeToString(bs)
}
