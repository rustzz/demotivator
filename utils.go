package demotivator

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func LoadSrcImage(path string) (imageFile *os.File) {
	imageFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	return imageFile
}

func LoadSrcImageFromURL(url string) (imageReader *bytes.Reader, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	imageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	imageReader = bytes.NewReader(imageBytes)
	return imageReader, nil
}
