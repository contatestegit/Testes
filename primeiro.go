package main

import (
	"fmt"
	"net/http"
)

// ///package main

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"os"

// 	"github.com/kr/pretty"
// 	"gitlab.b2w/team/alpha/omega.git/aggregation/sitemanager"
// )

// type AnywayReader struct {
// 	reader io.Reader
// }

// func (r *AnywayReader) FetchContent(context.Context) (content []byte, err error) {
// 	return ioutil.ReadAll(r.reader)
// }

// func NewAnywayReader(r io.Reader) *AnywayReader {
// 	return &AnywayReader{r}
// }

func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	uName, email, pwd, pwdConfirm := "", "", "", ""

	mux := http.NewServeMux()

	// Signup
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		/*
		   // Available for testing.
		   for key, value := range r.Form {
		       fmt.Printf("%s = %s\n", key, value)
		   }
		*/

		uName = r.FormValue("username")     // Data from the form
		email = r.FormValue("email")        // Data from the form
		pwd = r.FormValue("password")       // Data from the form
		pwdConfirm = r.FormValue("confirm") // Data from the form

		// Empty data checking
		uNameCheck := IsEmpty(uName)
		emailCheck := IsEmpty(email)
		pwdCheck := IsEmpty(pwd)
		pwdConfirmCheck := IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
			return
		}

		if pwd == pwdConfirm {
			// Save to database (username, email and password)
			fmt.Fprintln(w, "Registration successful.")
		} else {
			fmt.Fprintln(w, "Password information must be the same.")
		}
	})

	// Login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")  // Data from the form
		pwd = r.FormValue("password") // Data from the form

		// Empty data checking
		emailCheck := IsEmpty(email)
		pwdCheck := IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
			return
		}

		dbPwd := "1234!*."                   // DB simulation
		dbEmail := "cihan.ozhan@hotmail.com" // DB simulation

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintln(w, "Login succesful!")
		} else {
			fmt.Fprintln(w, "Login failed!")
		}
	})
	http.ListenAndServe(":8080", mux)
}
