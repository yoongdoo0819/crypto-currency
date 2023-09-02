package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/nomadcoders/nomadcoin/explorer"
	"github.com/nomadcoders/nomadcoin/rest"
)

func usage() {
	fmt.Println("Welcome !!")
	fmt.Println("Please use the following flags")
	fmt.Println("-port: Set the PORT of the server")
	fmt.Println("-mode: Choose between 'html' and 'rest'")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
