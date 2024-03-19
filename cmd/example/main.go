package main

import (
	"flag"
	"io"
	"os"
	"strings"

	lab2 "github.com/DimasComp/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "File with expression to compute")
	outputFile = flag.String("o", "", "File to write results")
	
	reader io.Reader
	writer io.Writer
)

func main() {
	flag.Parse()

	if (*inputExpression == "") == (*inputFile == "") {
		panic("Please specify exactly one of -e or -f")
	}

	if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		reader = file
	} else {
		reader = strings.NewReader(*inputExpression)
	}

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}
	handler := &lab2.ComputeHandler{
		Input: reader,
	    Output: writer,
	}
	err := handler.Compute()

	if err != nil {
		panic(err)
	}
}
