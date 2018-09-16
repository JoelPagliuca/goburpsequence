package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Items exported from Burp Proxy History
type Items struct {
	XMLName xml.Name `xml:"items"`
	Items   []Item   `xml:"item"`
}

// Item a HTTP request/response
type Item struct {
	XMLName xml.Name `xml:"item"`
	Host    string   `xml:"host"`
	Method  string   `xml:"method"`
	Status  int      `xml:"status"`
	Path    string   `xml:"path"`
}

// ReadBurpFile parses a burp file
func ReadBurpFile(fileLocation string) Items {
	xmlFile, err := os.Open(fileLocation)

	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	xmlFileBytes, _ := ioutil.ReadAll(xmlFile)

	var items Items
	xml.Unmarshal(xmlFileBytes, &items)

	return items
}
