package main

import (
	"flag"
	"fmt"
	lab2 "github.com/Makov-Vik/postfix-to-infix.git"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	fileInput       = flag.String("f", "", "File with expression to compute")
	fileOutput      = flag.String("o", "", "File to expression output")
)

func main() {
	flag.Parse()

	var Input io.Reader
	var Output io.Writer

	if len(*fileInput) > 0 {
		file, err := os.Open(*fileInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Such file does not exist.")
			return
		}
		Input = file
	} else {
		Input = strings.NewReader(*inputExpression)
	}

	if len(*fileOutput) > 0 {
		file, err := os.OpenFile(*fileOutput, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Such file does not exist.")
			return
		}
		Output = file
	} else {
		Output = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  Input,
		Output: Output,
	}
	err := handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
}
