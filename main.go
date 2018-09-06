package main

import (
	"flag"
	"fmt"
)

func usage() {
	fmt.Println("goburpsequence!")
	fmt.Println("")
	flag.PrintDefaults()
}

func main() {
	helpFlag := flag.Bool("h", false, "Display this help text")
	plantServer := flag.String("server", "http://www.plantuml.com/plantuml/uml/", "server to render the plantuml document")
	burpStateFile := flag.String("b", "", "burpstate file to convert")
	pcapFile := flag.String("p", "", "burpstate file to convert")
	detailLevel := flag.Int("d", 1, "level of detail for the diagram {1|2|3}")
	flag.Parse()

	if *helpFlag {
		usage()
	}

	fmt.Println("plantServer: " + *plantServer)
	fmt.Println("burpStateFile: " + *burpStateFile)
	fmt.Println("pcapFile: " + *pcapFile)
	fmt.Println("detailLevel: " + string(*detailLevel))
}
