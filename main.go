package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Welcome !!")
	fmt.Println("Please use the following commands")
	fmt.Println("explorer: Start the HTML Explorer")
	fmt.Println("rest: Start the REST API")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)
	portFlag := rest.Int("port", 4000, "Sets the port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(*portFlag)
		fmt.Println("Start Rest API")
	}

}
