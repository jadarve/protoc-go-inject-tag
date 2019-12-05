package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var inputFile string
	var outputFile string
	var xxxTags string
	flag.StringVar(&inputFile, "input", "", "path to input file")
	flag.StringVar(&outputFile, "output", "", "path to output file")
	flag.StringVar(&xxxTags, "XXX_skip", "", "skip tags to inject on XXX fields")

	flag.Parse()

	var xxxSkipSlice []string
	if len(xxxTags) > 0 {
		xxxSkipSlice = strings.Split(xxxTags, ",")
	}

	if len(inputFile) == 0 {
		log.Fatal("input file is mandatory")
	}

	if len(outputFile) == 0 {
		log.Fatal("output file is mandatory")
	}

	// copy input to output
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(outputFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	areas, err := parseFile(outputFile, xxxSkipSlice)
	if err != nil {
		log.Fatal(err)
	}
	if err = writeFile(outputFile, areas); err != nil {
		log.Fatal(err)
	}
}
