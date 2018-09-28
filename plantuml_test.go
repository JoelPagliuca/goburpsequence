package main

import "testing"
import "fmt"
import "net/http"

func TestToPlantURL(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Covert diagram to URL", args{"./test/data/diagram.uml"}, "SoWkIImgAStDuNBAJrBGjLDmpCbCJbMmKiX8pSd9vmBpGC90MQ1WUNwvZeAdGaPYMQf2ea9IVcLQKMPgNWgN0v00cW40"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileToPlantURL(tt.args.filename); got != tt.want {
				resp, _ := http.Get("http://www.plantuml.com/plantuml/uml/" + got)
				fmt.Println("Status:", resp.StatusCode)
				t.Errorf("ToPlantURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
