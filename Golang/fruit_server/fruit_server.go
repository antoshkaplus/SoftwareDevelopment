package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // underneath the package name is sqlite3.
	"gopkg.in/zeromq/goczmq.v4"
)

func show_all(db *sql.DB) {
	rows, err := db.Query("select * from fruit;")
	check_err(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var created string
		err = rows.Scan(&id, &name, &created)
		check_err(err)
		fmt.Println(id, name, created)
	}
	err = rows.Err()
	check_err(err)
}

func make_tables(db *sql.DB) {
	// created is a string datetime.
	qry := `
	CREATE TABLE IF NOT EXISTS fruit (
		id integer primary key AUTOINCREMENT, 
		name text,
		created text);
	`
	_, err := db.Exec(qry)
	if err != nil {
		log.Fatalf("%q: %s\n", err, qry)
	}
}

func add_fruit(db *sql.DB, name string) {
	fmt.Println("Add")
	stmt, err := db.Prepare("insert into fruit(name, created) values(?, ?)")
	check_err(err)
	defer stmt.Close()

	_, err = stmt.Exec(name, time.Now().Format(time.RFC3339))
	check_err(err)
}

func delete_fruit(db *sql.DB, name string) {
	fmt.Println("Delete")
	stmt, err := db.Prepare("delete from fruit where name=?;")
	check_err(err)
	defer stmt.Close()

	_, err = stmt.Exec(name)
	check_err(err)
}

func handle_server() {
	db, err := sql.Open("sqlite3", "../tmp/fruit.db")
	check_err(err)
	defer db.Close()
	make_tables(db)

	sock, err := goczmq.NewRep("tcp://127.0.0.1:5555")
	check_err(err)
	defer sock.Destroy()

	for {
		fmt.Println("Ready to receive.")
		bytes, _, err := sock.RecvFrame()
		check_err(err)

		msg := map[string]string{}
		err = json.Unmarshal(bytes, &msg)
		check_err(err)

		fmt.Printf("Recv: %v\n", msg)

		res := "Success"
		switch msg["type"] {
		case "add":
			add_fruit(db, msg["name"])
		case "list":
			show_all(db)
		case "del":
			delete_fruit(db, msg["name"])
		default:
			res = "Failure"
		}

		msg_res, err := json.Marshal(res)
		check_err(err)
		err = sock.SendFrame(msg_res, 0)
		check_err(err)
	}
}
