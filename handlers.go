package main

import (
	"fmt"
	"net/http"
	"github.com/futurenda/google-auth-id-token-verifier"
	ct "github.com/cvhariharan/Utils/customtype"
	"github.com/cvhariharan/Utils/utils"
	"os"
	"encoding/json"
)

// signupHandler requires atleast fname, lname, username, password, email
func signupHandler(w http.ResponseWriter, r *http.Request) {
	var jwt string
	var user ct.User
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	fname := r.Form.Get("fname")
	lname := r.Form.Get("lname")
	email := r.Form.Get("email")
	googleAuth := r.Form.Get("googleauth")
	if username != "" && password != "" && fname != "" && lname != "" && email != "" {
		user = ct.User{
			FName: fname,
			LName: lname,
			UName: username,
			Email: email,
		}
		if googleAuth == "true" {
			// Nothing changes much, could be done with just a single line but here it goes
			// If googleAuth is enabled, password is the google id token
			v := googleAuthIDTokenVerifier.Verifier{}
			err := v.VerifyIDToken(password, []string{
				os.Getenv("CLIENTID"),
			})
			if err != nil {
				claimSet, err := googleAuthIDTokenVerifier.Decode(password)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(claimSet)
				if email != claimSet.Email {
					// If the email id does not match the token email
					fmt.Fprint(w, `{ "error": "Not a valid email id" }`)
					return
				}
				// Only adds this extra profile pic if signed up with google
				user.ProfilePic = ct.Image{Link: claimSet.Picture}
			} else {
				fmt.Println(err)
			}
		}
		user.CreatePassword(password)
		jwt = utils.UserSignup(user, session)
		if jwt == "" {
			fmt.Fprint(w, `{ "error": "username or email exists" }`)
		} else {
			fmt.Fprint(w, `{ "result": "success", "token": "` + jwt + "\"}")
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
	googleAuth := r.Form.Get("googleauth")

	if googleAuth == "true" {
		// If googleauth is enabled, username (email) is extracted from a valid google id token
		v := googleAuthIDTokenVerifier.Verifier{}
		err := v.VerifyIDToken(password, []string{
			os.Getenv("CLIENTID"),
		})
		if err != nil {
			claimSet, err := googleAuthIDTokenVerifier.Decode(password)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(claimSet)
			u := utils.GetUser(username, session)
			fmt.Println(u)
			if u.Email != claimSet.Email {
				// If the email id does not match the token email
				fmt.Fprint(w, `{ "error": "Not a valid email id" }`)
				return
			}
			
		}
	}
	if username != "" && password != "" {
		jwt := utils.UserLogin(username, password, session)
		if jwt == "" {
			fmt.Fprint(w, `{ "error": "could not login" }`)
		} else {
			fmt.Fprint(w, `{ "result": "success", "token": "` + jwt + "\"}")
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
	jwt := utils.GenerateJWT(follower, session)
	jsonString += jwt + "\" }"
	w.Write([]byte(jsonString))
}

func unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var jsonString string
	followee := r.Form.Get("followee")
	follower := r.Form.Get("username")
	w.Header().Set("Content-Type", "application/json")
	if utils.UnfollowUser(follower, followee, session) {
		jsonString = `{ "result": "success", "token" : "`
	} else {
		jsonString = `{ "error": "could not unfollow", "token" : "`
	}
	jwt := utils.GenerateJWT(follower, session)
	jsonString += jwt + "\" }"
	w.Write([]byte(jsonString))
}

func getprofile(w http.ResponseWriter, r *http.Request) {
	var result interface{}
	var jsonString string
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	profilename := r.Form.Get("profilename")
	username := r.Form.Get("username")

	result = utils.GetProfile(profilename, session)
	if result == nil {
		// If res is empty
		jsonString = `{ "error": "could not locate the profile", "token" : "`
	} else {
		j, _ := json.Marshal(result)
		jsonString = `{ "result": ` + string(j) + `, "token": "` + utils.GenerateJWT(username, session) + "\"}"
	}

	w.Write([]byte(jsonString))
}

// updateDetails
// func updateDetails(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	username := r.Form.Get("username")
// 	password := r.Form.Get("password")
// 	newUsername := r.Form.Get("newUsername")
// 	newPassword := r.Form.Get("newPassword")
// 	fname := r.Form.Get("fname")
// 	lname := r.Form.Get("lname")
// 	email := r.Form.Get("email")
// 	// profilepic := r.Form.Get("image")

// 	user := ct.User{}

// 	user.UName = newUsername
// 	user.FName = fname
// 	user.LName = lname
// 	user.Email = email
// 	user.CreatePassword(newPassword)

// 	conf := utils.UpdateDetails(username, user, session)
// 	if conf == "" {
// 		fmt.Fprint(w, `{"result": "success"}`)
// 	} else {
// 		fmt.Fprint(w, `{"result": "error"}`)
// 	}

// }
