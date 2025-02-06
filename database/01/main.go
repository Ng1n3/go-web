package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	dsn := "myuser:00005@tcp(127.0.0.1:3306)/aws?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("unable to reach the database: ", err)
	}

	log.Println("Connected to the database successfully!")

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":3050", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "At Index")
	checkError(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT name FROM amigos;`)
	checkError(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		checkError(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE amigos (name VARCHAR(20));`)
	checkError(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	checkError(err)

	n, err := r.RowsAffected()
	checkError(err)

	fmt.Println(w, "CREATED TABLE amigos", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO amigos VALUES("James");`)
	checkError(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	checkError(err)

	n, err := r.RowsAffected()
	checkError(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM amigos;`)
	checkError(err)
	defer rows.Close()

	var name string
	var id int
	var createdAt string
	for rows.Next() {
		err = rows.Scan(&id, &name, &createdAt)
		checkError(err)
		fmt.Fprintf(w, "ID: %d, Name: %s, Created At: %s\n", id, name, createdAt)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(` UPDATE amigos SET name="Jimmy" where name="james";`)
	checkError(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	checkError(err)

	n, err := r.RowsAffected()
	checkError(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM amigos WHERE name="Jimmy";`)
	checkError(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	checkError(err)

	n, err := r.RowsAffected()
	checkError(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE amigos;`)
	checkError(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	checkError(err)

	fmt.Fprintln(w, "DROPPED TABLE amigos")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
