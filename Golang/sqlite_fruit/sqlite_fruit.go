package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// In Go, the underscore _ in an import statement is known as a blank identifier.
	// It's used to import a package solely for its side effects, without explicitly
	// using any of its exported functions or variables. This is particularly useful
	// for packages that perform initialization tasks in their init functions.
	// When a package is imported with an underscore, its init functions are executed,
	// which can register database drivers, set up configurations, or perform other setup tasks.
	// However, the package's namespace is not imported, meaning its functions and
	// variables cannot be directly accessed.

	_ "github.com/mattn/go-sqlite3" // underneath the package name is sqlite3.
)

func add() {

}

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

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	check_err(err)

	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed
	// until the surrounding function returns.
	// Called even in case panic.
	defer db.Close()

	// created is a string datetime.
	qry := `
		create table fruit (
			id integer primary key AUTOINCREMENT, 
			name text,
			created text);
	`
	_, err = db.Exec(qry)
	if err != nil {
		log.Fatalf("%q: %s\n", err, qry)
	}

	stmt, err := db.Prepare("insert into fruit(name, created) values(?, ?)")
	check_err(err)
	defer stmt.Close()

	_, err = stmt.Exec("Apple", time.Now().Format(time.RFC3339))
	check_err(err)

	fmt.Println("Print all after insert.")
	show_all(db)

	db.Exec("delete from fruit where name = 'Apple';")

	fmt.Println("Print all after remove.")
	show_all(db)

	log.Println("End Program")
}
