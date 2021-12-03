package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Printf("hit the home page endpoint")
}
func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	fmt.Println("Hit the users page endpoint")
	json.NewEncoder(w).Encode(users)
}

func getUsers() []*User {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	if err != nil {

		panic(err)
	}
	fmt.Println(db)
	defer db.Close()

	// Execute query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var user []*User

	for results.Next() {
		var usr User
		err = results.Scan(&usr.ID, &usr.Name)
		if err != nil {
			panic(err)
		}
		user = append(user, &usr)
	}

	return user

}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
