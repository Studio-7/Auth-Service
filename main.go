package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	ct "github.com/cvhariharan/Data-Models/customtype"
	"github.com/cvhariharan/Utils/utils"
	"github.com/joho/godotenv"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

var session *r.Session

// signupHandler requires atleast fname, lname, username, password
func signupHandler(w http.ResponseWriter, r *http.Request) {
	data := decodeRequest(r.Body)
	username, userOk := data["username"]
	password, passOk := data["password"]
	if userOk && passOk {
		user := ct.User{
			FName: data["fname"].(string),
			LName: data["lname"].(string),
			UName: username.(string),
		}
		user.CreatePassword(password.(string))
		jwt := utils.UserSignup(user, session)
		if jwt == "" {
			fmt.Fprint(w, "Error: Username exists")
		} else {
			fmt.Fprint(w, jwt)
		}
	} else {
		fmt.Fprint(w, "Error: Check request params")
	}
}

// loginHandler requires username and password
func loginHandler(w http.ResponseWriter, r *http.Request) {
	data := decodeRequest(r.Body)
	username, userOk := data["username"].(string)
	password, passOk := data["password"].(string)
	if userOk && passOk {
		jwt := utils.UserLogin(username, password, session)
		if jwt == "" {
			fmt.Fprint(w, "Error: Could not authenticate the user. Check username and password or try again later.")
		} else {
			fmt.Fprint(w, jwt)
		}
	} else {
		fmt.Fprint(w, "Error: Check request params")
	}
}

// decodeRequest takes in request body and returns the map
func decodeRequest(request io.ReadCloser) map[string]interface{} {
	var data map[string]interface{}
	req, err := ioutil.ReadAll(request)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(req), &data)
	return data
}

func main() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}

	url := os.Getenv("DBURL")
	s, err := r.Connect(r.ConnectOpts{
		Address: url, // endpoint without http
	})
	if err != nil {
		log.Fatalln(err)
	}
	session = s
	port := ":" + os.Getenv("PORT")
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
