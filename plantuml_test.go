package main

import "testing"

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
			if got := ToPlantURL(tt.args.filename); got != tt.want {
				t.Errorf("ToPlantURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
