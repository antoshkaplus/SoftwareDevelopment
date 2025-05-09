package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/akamensky/argparse"
	_ "github.com/mattn/go-sqlite3" // underneath the package name is sqlite3.
)

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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

func main() {
	db, err := sql.Open("sqlite3", "../tmp/fruit.db")
	check_err(err)
	defer db.Close()
	make_tables(db)

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

	err = parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		return
	}

	if add_cmd.Happened() {
		fmt.Println("Add")
		stmt, err := db.Prepare("insert into fruit(name, created) values(?, ?)")
		check_err(err)
		defer stmt.Close()

		_, err = stmt.Exec(add_fruit_name, time.Now().Format(time.RFC3339))
		check_err(err)
	}
	if list_cmd.Happened() {
		fmt.Println("List")
		show_all(db)
	}
	if del_cmd.Happened() {
		fmt.Println("Delete")
		stmt, err := db.Prepare("delete from fruit where name=?;")
		check_err(err)
		defer stmt.Close()

		_, err = stmt.Exec(del_fruit_name)
		check_err(err)
	}

	log.Println("End Program")
}
