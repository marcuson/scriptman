package handle

import (
	"marcuson/scriptman/internal/script/internal/processor"
	"os"
	"testing"

	"github.com/fluentassert/verify"
	"github.com/stretchr/testify/mock"
)

const (
	testdir = "../../_testdata"
)

func TestHandlerPassthroughOk(t *testing.T) {
	file, err := os.Open(testdir + "/meta_ok_001.sh")
	if err != nil {
		t.Fatal("unable to open test file")
	}
	defer file.Close()

	mockProc := processor.NewMockProcessor(t)
	mockProc.EXPECT().IsProcessCompletedEarly().Return(false)
	mockProc.EXPECT().ProcessStart().Return(nil).Times(1)
	mockProc.EXPECT().ProcessEnd().Return(nil).Times(1)
	mockProc.EXPECT().
		ProcessLine(mock.AnythingOfType("*scan.LineScript")).
		Return(nil).
		Times(7)
	handler := NewHandler(file, mockProc)

	err = handler.Handle()
	verify.NoError(err).Require(t)
	verify.String(handler.Interpreter()).Equal("bash").Require(t)
}
