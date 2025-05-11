package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/akamensky/argparse"
	"gopkg.in/zeromq/goczmq.v4"
)

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

		sock, err := goczmq.NewReq("tcp://127.0.0.1:5555")
		check_err(err)
		defer sock.Destroy()

		msg := map[string]string{"type": "add", "name": *add_fruit_name}
		bytes, err := json.Marshal(msg)
		check_err(err)

		err = sock.SendFrame(bytes, 0)
		check_err(err)
		fmt.Printf("Message Sent: %v\n", bytes)
		time.Sleep(5 * time.Second)
		// stmt, err := db.Prepare("insert into fruit(name, created) values(?, ?)")
		// check_err(err)
		// defer stmt.Close()

		// _, err = stmt.Exec(add_fruit_name, time.Now().Format(time.RFC3339))
		// check_err(err)
	}
	if list_cmd.Happened() {
		fmt.Println("List")
		// show_all(db)
	}
	if del_cmd.Happened() {
		fmt.Printf("Delete %v\n", *del_fruit_name)
		// stmt, err := db.Prepare("delete from fruit where name=?;")
		// check_err(err)
		// defer stmt.Close()

		// _, err = stmt.Exec(del_fruit_name)
		// check_err(err)
	}

}
