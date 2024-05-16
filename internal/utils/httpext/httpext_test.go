package httpext

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/fluentassert/verify"
	"github.com/prashantv/gostub"
)

type fakeHttpBody struct {
	buf *bytes.Buffer
}

func newFakeHttpBody(content string) *fakeHttpBody {
	buf := new(bytes.Buffer)
	buf.WriteString(content)
	return &fakeHttpBody{
		buf: buf,
	}
}

func (obj *fakeHttpBody) Read(p []byte) (int, error) {
	return obj.buf.Read(p)
}

func (obj *fakeHttpBody) Close() error {
	return nil
}

func TestDownloadFileOk(t *testing.T) {
	fakeResp := &http.Response{
		StatusCode: 200,
		Body:       newFakeHttpBody("test"),
	}
	stub := gostub.StubFunc(&httpGet, fakeResp, nil)
	defer stub.Reset()

	out := new(bytes.Buffer)
	err := DownloadFile("http://example.com", out)
	verify.NoError(err).Assert(t)

	res := out.String()
	verify.String(res).Equal("test").Require(t)
}

func TestDownloadFileNot200(t *testing.T) {
	fakeResp := &http.Response{
		StatusCode: 404,
		Body:       newFakeHttpBody("not found"),
	}
	stub := gostub.StubFunc(&httpGet, fakeResp, nil)
	defer stub.Reset()

	out := new(bytes.Buffer)
	err := DownloadFile("http://example.com", out)
	verify.Error(err).Contain("download failed").Assert(t)
}
