package main

import (
	"fmt"
	"github.com/ashinsabu/harness-tool/cmd"
	"os"
)

var version = "0.1.0"

func main() {
	if err := cmd.Execute(version); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
