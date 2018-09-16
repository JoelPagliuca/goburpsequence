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
	burpItemsFile := flag.String("b", "", "burp saved items file to convert")
	pcapFile := flag.String("p", "", "burpstate file to convert")
	detailLevel := flag.Int("d", 1, "level of detail for the diagram {1|2|3}")
	outputPNG := flag.String("output-png", "diagram.png", "output a PNG file (default)")
	outputSVG := flag.String("output-svg", "diagram.svg", "output a SVG file")
	flag.Parse()

	if *helpFlag {
		usage()
	}

	fmt.Println("plantServer: " + *plantServer)
	fmt.Println("burpStateFile: " + *burpItemsFile)
	fmt.Println("pcapFile: " + *pcapFile)
	fmt.Println("detailLevel: " + string(*detailLevel))
	fmt.Println("outputPNG: " + *outputPNG)
	fmt.Println("outputSVG: " + *outputSVG)
}
