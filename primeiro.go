package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"golang.org/x/net/publicsuffix"
)

// const (
// 	baseURL = "http://admsite.b2w"
// )

// type App struct {
// 	Client *http.Client
// }

// type AuthenticationCookies struct {
// 	Cookies []*http.Cookie
// }

// func (app *App) Login(username string, password string) error {
// 	client := app.Client

// 	loginFormURL := baseURL + "/SiteManagerWeb/login.jsp"

// 	resp, err := client.Get(loginFormURL)

// 	if err != nil {
// 		log.Fatalln(err)
// 		return nil
// 	}

// 	data := url.Values{
// 		"j_username": {username},
// 		"j_password": {password},
// 	}

// 	loginPostURL := baseURL + "/SiteManagerWeb/j_security_check"
// 	resp, err = client.PostForm(loginPostURL, data)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		log.Fatalln(errors.Wrap(errors.New("Erro "+strconv.Itoa(resp.StatusCode)+" na requisição"), "bla"))
// 	}

// 	defer resp.Body.Close()
// 	return nil
// }

func main() {

	cj, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})

	if err != nil {
		log.Fatalln(err)
	}

	app := App{
		Client: &http.Client{Jar: cj},
	}
	username := "wedney.lima.Acom"
	password := "Acom#1234"

	app.login(username, password)

	u, err := url.Parse("http://admsite.b2w/SiteManagerWeb/")
	ac := AuthenticationCookies{
		Cookies: cj.Cookies(u),
	}
	fmt.Println(ac.Cookies)

	req, err := http.NewRequest("GET", "http://admsite.b2w/SiteManagerWeb/menu/menuXML.do?id=203682", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set(ac.Cookies[0].Name, ac.Cookies[0].Value)
	resp, err := app.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	outFile, err := os.Create("output.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
