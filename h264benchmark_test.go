package h264benchmark

import (
	"github.com/pion/webrtc/pkg/media/h264reader"
	"io"
	"os"
	"testing"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestCurrentH264Reader(t *testing.T) {
	file, err := os.Open("test.h264")
	checkErr(err)
	reader, err := h264reader.NewReader(file)
	checkErr(err)
	for {
		_, err := reader.NextNAL()
		if err == io.EOF {
			break
		}
		checkErr(err)
	}
}

func TestNewH264Reader(t *testing.T) {
	file, err := os.Open("test.h264")
	checkErr(err)
	newReader, err := NewReader(file)
	checkErr(err)
	for {
		_, err := newReader.NextNAL()
		if err == io.EOF {
			break
		}
		checkErr(err)
	}
}
