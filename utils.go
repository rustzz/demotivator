package demotivator

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
)

func LoadSrcImage(path string) (imageFile *os.File, err error) {
	imageFile, err = os.Open(path)
	if err != nil {
		return
	}
	return
}

func LoadSrcImageFromURL(url string) (imageReader *bytes.Reader, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	imageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	imageReader = bytes.NewReader(imageBytes)
	return
}
