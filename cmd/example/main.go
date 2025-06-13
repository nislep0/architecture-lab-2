package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/nislep0/architecture-lab-2"
)

var (
	exprFlag   = flag.String("e", "", "Expression to compute")
	fileFlag   = flag.String("f", "", "Path to input file")
	outputFlag = flag.String("o", "", "Path to output file")
)

func main() {
	flag.Parse()

	if (*exprFlag != "" && *fileFlag != "") || (*exprFlag == "" && *fileFlag == "") {
		fmt.Fprintln(os.Stderr, "Error: must provide either -e or -f, but not both or neither.")
		os.Exit(1)
	}

	var inputReader io.Reader
	if *exprFlag != "" {
		inputReader = strings.NewReader(*exprFlag + "\n")
	} else {
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		inputReader = file
	}

	var outputWriter io.Writer
	if *outputFlag != "" {
		file, err := os.Create(*outputFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		outputWriter = file
	} else {
		outputWriter = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  inputReader,
		Output: outputWriter,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Computation error: %v\n", err)
		os.Exit(1)
	}
}