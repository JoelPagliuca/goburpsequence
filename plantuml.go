package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
)

//http://plantuml.com/code-javascript-synchronous

// ToPlantURL generate the URL from the diagram file
func ToPlantURL(filename string) string {
	// open the file
	diagramData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	deflatedData := deflate(diagramData)
	// encode
	return string(deflatedData[:])
}

func deflate(input []byte) []byte {
	var output bytes.Buffer
	writer, err := zlib.NewWriterLevel(&output, zlib.BestCompression)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer writer.Close()
	writer.Write(input)
	return output.Bytes()
}
