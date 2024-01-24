package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Welcome to Blockchain Service")
	fmt.Println()
	fmt.Println("Please use the following commands:")
	fmt.Println()
	fmt.Println("explorer:	Start the HTML Explorer")
	fmt.Println("rest:		Start the REST API (recommanded)")
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
		fmt.Println("Start REST API")
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(fmt.Sprint(`Start server `, *portFlag))
	}
}
