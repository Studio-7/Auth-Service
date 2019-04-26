package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/joho/godotenv"
	"github.com/cvhariharan/Utils/utils"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

var session *r.Session

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

	// Load only returns the error; nil if none.
	e := godotenv.Load()
	if e != nil {
		// Fatal-> Print+os.exit(1)
		log.Fatal(e)
	}
	endpoints := os.Getenv("DBURL")
	url := strings.Split(endpoints, ",")
	dbpass := os.Getenv("DBPASS")
	s, err := r.Connect(r.ConnectOpts{
		Addresses: url,
		Password:  dbpass,
	})
	if err != nil {
		log.Fatalln(err)
	}
	session = s
	port := ":" + os.Getenv("PORT")
	http.HandleFunc("/user/signup", signupHandler)
	http.HandleFunc("/user/login", loginHandler)
	http.HandleFunc("/user/followuser", utils.AuthMiddleware(followUserHandler, session))
	http.HandleFunc("/user/unfollowuser", utils.AuthMiddleware(unfollowUserHandler, session))
	http.HandleFunc("/user/getprofile", utils.AuthMiddleware(getprofile, session))
	// http.HandleFunc("/user/updatedetails", updateDetails)
	log.Fatal(http.ListenAndServe(port, nil))
}
