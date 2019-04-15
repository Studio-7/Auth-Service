package main

import (
	"fmt"
	"net/http"

	ct "github.com/cvhariharan/Utils/customtype"
	"github.com/cvhariharan/Utils/utils"
)

// signupHandler requires atleast fname, lname, username, password, email
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
			fmt.Fprint(w, `{ "error": "username exists" }`)
		} else {
			fmt.Fprint(w, jwt)
		}
	} else {
		fmt.Fprint(w, `{ "error": "check request params" }`)
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
			fmt.Fprint(w, `{ "error": "could not authenticate, check username or password and try again later" }`)
		} else {
			fmt.Fprint(w, jwt)
		}
	} else {
		fmt.Fprint(w, `{ "error": "check request params" }`)
	}
}

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

// updateDetails
func updateDetails(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	newUsername := r.Form.Get("newUsername")
	newPassword := r.Form.Get("newPassword")
	fname := r.Form.Get("fname")
	lname := r.Form.Get("lname")
	email := r.Form.Get("email")
	// profilepic := r.Form.Get("image")

	user := ct.User{}

	user.UName = newUsername
	user.FName = fname
	user.LName = lname
	user.Email = email
	user.CreatePassword(newPassword)

	conf := utils.UpdateDetails(username, user, session)
	if conf == "" {
		fmt.Fprint(w, `{"result": "success"}`)
	} else {
		fmt.Fprint(w, `{"result": "error"}`)
	}

}
