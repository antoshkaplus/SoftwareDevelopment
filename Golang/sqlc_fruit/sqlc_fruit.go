package main

import (
	"context"
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
	fruitdb "fruit/db"

	_ "github.com/mattn/go-sqlite3" // underneath the package name is sqlite3.
)

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func make_sql_str(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func show_all(fruit_list []fruitdb.Fruit) {
	for _, value := range fruit_list {
		_, err := fmt.Printf("%v %v %v\n", value.ID, value.Name.String, value.Created.String)
		check_err(err)
	}
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

	ctx := context.Background()
	mydb := fruitdb.New(db)

	params := fruitdb.CreateFruitParams{
		Name:    make_sql_str("Apple"),
		Created: make_sql_str(time.Now().Format(time.RFC3339)),
	}
	_, err = mydb.CreateFruit(ctx, params)
	check_err(err)

	fmt.Println("Print all after insert.")
	fruit_list, err := mydb.ListFruit(ctx)
	check_err(err)
	show_all(fruit_list)

	mydb.DeleteFruit(ctx, make_sql_str("Apple"))

	fmt.Println("Print all after remove.")
	fruit_list, err = mydb.ListFruit(ctx)
	check_err(err)
	show_all(fruit_list)

	log.Println("End Program")
}
