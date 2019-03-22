package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	ct "github.com/cvhariharan/Utils/customtype"
	"github.com/cvhariharan/Utils/utils"
	"github.com/joho/godotenv"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

var session *r.Session

// signupHandler requires atleast fname, lname, username, password
func signupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	fname := r.Form.Get("fname")
	lname := r.Form.Get("lname")
	email := r.Form.Get("email")
	if username != "" && password != "" && fname != "" && lname != "" && email != "" {
		user := ct.User{
			FName: fname,
			LName: lname,
			UName: username,
			Email: email,
		}
		user.CreatePassword(password)
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
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username != "" && password != "" {
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
	endpoints := os.Getenv("DBURL")
	url := strings.Split(endpoints, ",")
	s, err := r.Connect(r.ConnectOpts{
		Addresses: url, // endpoint without http
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
