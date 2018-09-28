package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
)

const asciiSpace = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

// FileToPlantURL generate the URL from the diagram file
func FileToPlantURL(filename string) string {
	// open the file
	diagramData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	deflatedData := deflate(diagramData)
	encodedData := encode(deflatedData)
	return string(encodedData)
}

func deflate(input []byte) []byte {
	var output bytes.Buffer
	writer := zlib.NewWriter(&output)
	defer writer.Close()
	writer.Write(input)
	writer.Flush()
	return output.Bytes()
}

// encode ASCII transformation _close_ to base64
func encode(input []byte) string {
	var output bytes.Buffer
	length := len(input)
	for i := 0; i < 3-length%3; i++ {
		input = append(input, byte(0))
	}

	for i := 0; i < length; i += 3 {
		b1, b2, b3, b4 := input[i], input[i+1], input[i+2], byte(0)

		b4 = b3 & 0x3f
		b3 = ((b2 & 0xf) << 2) | (b3 >> 6)
		b2 = ((b1 & 0x3) << 4) | (b2 >> 4)
		b1 = b1 >> 2

		for _, b := range []byte{b1, b2, b3, b4} {
			output.WriteByte(byte(asciiSpace[b]))
		}
	}
	return string(output.Bytes())
}
