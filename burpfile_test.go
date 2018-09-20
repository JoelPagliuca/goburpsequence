package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestReadBurpFile(t *testing.T) {
	fmt.Println("\nItems:")
	items := ReadBurpFile("./test/data/items.xml")

	for idx, item := range items.Items {
		t.Run(fmt.Sprintf("item %v has a valid hostname", idx), func(t *testing.T) {
			if !(item.Host == "joelpagliuca.me" || item.Host == "fonts.gstatic.com") {
				t.Errorf("item.Host = %v ", item.Host)
			}
		})

		t.Run(fmt.Sprintf("item %v has a valid path", idx), func(t *testing.T) {
			match, _ := regexp.MatchString(`^/.*`, item.Path)
			if !match {
				t.Errorf("item.Path = %v", item.Path)
			}
		})

		t.Run(fmt.Sprintf("item %v has the correct method", idx), func(t *testing.T) {
			if item.Method != "GET" {
				t.Errorf("item.Method = %v", item.Method)
			}
		})

		t.Run(fmt.Sprintf("item %v has a valid response code", idx), func(t *testing.T) {
			if item.Status != 200 && item.Status != 404 {
				t.Errorf("item.Host = %v ", item.Status)
			}
		})
	}
}
