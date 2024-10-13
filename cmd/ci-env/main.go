package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var format string
	flag.StringVar(&format, "format", "json", "output format (json or export)")
	flag.Parse()

	if format == "" {
		flag.Usage()
		return
	}
	err := PrintPlatformInfo(format)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
