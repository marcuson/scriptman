package hashext

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestMd5Str(t *testing.T) {
	in := "example"
	res := Md5Str(in)
	verify.String(res).Equal("1a79a4d60de6718e8e5b326e338ae533").Require(t)
}

func TestSha256Str(t *testing.T) {
	in := "example"
	res := Sha256Str(in)
	verify.String(res).Equal("50d858e0985ecc7f60418aaf0cc5ab587f42c2570a884095a9e8ccacd0f6545c").
		Require(t)
}
