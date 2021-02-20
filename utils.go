package demotivator

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func LoadSrcImage(path string) *os.File {
	imageReader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return imageReader
}

func LoadSrcImageFromURL(url string) bytes.Reader {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return bytes.Reader{}
	}
	defer resp.Body.Close()
	imageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return bytes.Reader{}
	}
	imageReader := bytes.NewReader(imageBytes)
	return *imageReader
}
