package main

import (
	"github.com/yuntan/v4l2/capture"

	"bytes"
	"fmt"
	"image/jpeg"
	"io/ioutil"
)

const (
	VIDEO_DEVICE = "/dev/video0"
)

func main() {
	img, err := capture.Capture(VIDEO_DEVICE)
	if err != nil {
		fmt.Println("Error: failed to capture: ", err)
		return
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile("test.jpg", buf.Bytes(), 0644)
}
