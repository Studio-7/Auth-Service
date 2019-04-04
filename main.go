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

func followUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var jsonString string
	followee := r.Form.Get("followee")
	follower := r.Form.Get("username")
	w.Header().Set("Content-Type", "application/json")
	if utils.FollowUser(follower, followee, session) {
		jsonString = `{ "result": "success", "token" : "`
	} else {
		jsonString = `{ "error": "could not follow", "token" : "`
	}
	jwt := utils.GenerateJWT(followee, session)
	jsonString += jwt + "\" }"
	w.Write([]byte(jsonString))
}

// signupHandler requires atleast fname, lname, username, password, email
func signupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	fname := r.Form.Get("fname")
	lname := r.Form.Get("lname")
	email := r.Form.Get("email")
	w.Header().Set("Content-Type", "application/json")
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
			fmt.Fprint(w, `{ "error": "username or email exists" }`)
		} else {
			fmt.Fprint(w, `{ "result": "success", "jwt": "` + jwt + "\"}")
		}
	} else {
		fmt.Fprint(w, `{ "error": "check request params" }`)
	}
}

// loginHandler requires username and password
func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username != "" && password != "" {
		jwt := utils.UserLogin(username, password, session)
		if jwt == "" {
			fmt.Fprint(w, `{ "error": "could not authenticate, check username or password and try again later" }`)
		} else {
			fmt.Fprint(w, `{ "result": "success", "jwt": "` + jwt + "\"}")
		}
	} else {
		fmt.Fprint(w, `{ "error": "check request params" }`)
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
	dbpass := os.Getenv("DBPASS")
	s, err := r.Connect(r.ConnectOpts{
		Addresses: url,
		Password: dbpass,
	})
	if err != nil {
		log.Fatalln(err)
	}
	session = s
	port := ":" + os.Getenv("PORT")
	http.HandleFunc("/user/signup", signupHandler)
	http.HandleFunc("/user/login", loginHandler)
	http.HandleFunc("/user/followuser", utils.AuthMiddleware(followUserHandler, session))
	log.Fatal(http.ListenAndServe(port, nil))
}
