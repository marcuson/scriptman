package httpext

import (
	"fmt"
	"io"
	"net/http"
)

var (
	httpGet = http.Get
)

func DownloadFile(url string, writer io.Writer) error {
	resp, err := httpGet(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("download failed, status code is not 200")
	}

	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
