package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	messag string "json:message"
}
type User struct {
	ID        int    "json:id"
	FirstName string "json:firstname"
	LastName  string "json:lastname"
	Age       int    "json:age"
}

func main() {
	apiRoot := "/api"

	// ...// api
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := User{
			ID:        123,
			FirstName: "ahmet can",
			LastName:  "ceylan",
			Age:       23,
		}
		output, err := json.Marshal(message)
		checkerror(err)
		//w.Header().Set("content.type","application/json") // default olarak kullanildigi icin yazmak zorunda degilsinhiz
		fmt.Fprintf(w, string(output))
	})

	// .../api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{ID: 1, FirstName: "example", LastName: "example", Age: 23},
			User{ID: 2, FirstName: "example1", LastName: "example1", Age: 22},
			User{ID: 3, FirstName: "example2", LastName: "example2", Age: 1},
		}
		message := users
		output, err := json.Marshal(message)
		checkerror(err)
		fmt.Fprintf(w, string(output))
	})

	// .../api/me
	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{1, "example", "example", 23}
		message := user
		output, err := json.Marshal(message)
		checkerror(err)
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8080", nil)

}

func checkerror(err error) {
	if err != nil {
		fmt.Println("fatal error : ", err.Error())
		os.Exit(1)
	}
}
