package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/akamensky/argparse"
	"gopkg.in/zeromq/goczmq.v4"
)

func send_msg(msg map[string]string) {
	sock, err := goczmq.NewReq("tcp://127.0.0.1:5555")
	check_err(err)
	defer sock.Destroy()

	bytes, err := json.Marshal(msg)
	check_err(err)

	err = sock.SendFrame(bytes, 0)
	check_err(err)
	fmt.Printf("Message Sent: %v\n", bytes)
	time.Sleep(10 * time.Second)
}

func handle_client_cli(args []string) {
	parser := argparse.NewParser("Fruit CLI", "Perform fruit db operations.")
	add_cmd := parser.NewCommand("add", "")
	var add_fruit_name *string = add_cmd.StringPositional(
		&argparse.Options{
			Required: true,
			Help:     "Fruit name to add."})
	list_cmd := parser.NewCommand("list", "")
	del_cmd := parser.NewCommand("del", "")
	var del_fruit_name *string = del_cmd.StringPositional(
		&argparse.Options{
			Required: true,
			Help:     "Fruit name to delete."})

	err := parser.Parse(args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		return
	}

	if add_cmd.Happened() {
		fmt.Printf("Add %v\n", *add_fruit_name)
		msg := map[string]string{"type": "add", "name": *add_fruit_name}
		send_msg(msg)
	}
	if list_cmd.Happened() {
		fmt.Println("List")
		msg := map[string]string{"type": "list"}
		send_msg(msg)
	}
	if del_cmd.Happened() {
		fmt.Printf("Delete %v\n", *del_fruit_name)
		msg := map[string]string{"type": "del", "name": *add_fruit_name}
		send_msg(msg)
	}

}
