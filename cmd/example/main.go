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
	inputExpression = flag.String("e", "17 10 + 3 * 9 /", "Expression to compute")
	fileInput       = flag.String("f", "", "File with expression to compute")
)

func main() {
	flag.Parse()

	var Input io.Reader

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
	//TODO work with output and with handler
}
