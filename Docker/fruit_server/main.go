package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser(
		"Fruit Client-Server",
		"Pick if want to be a server or to make client request.")
	client_cmd := parser.NewCommand("client", "Make client request")
	server_cmd := parser.NewCommand("server", "Start application as a server")

	err := parser.Parse(os.Args[:2])
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		return
	}

	if client_cmd.Happened() {
		args := slices.Delete(os.Args, 1, 2)
		handle_client_cli(args)
	}
	if server_cmd.Happened() {
		handle_server()
	}
}
