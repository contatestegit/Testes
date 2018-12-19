package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/net/publicsuffix"
)

const (
	baseURL = "http://admsite.b2w"
)

type App struct {
	Client *http.Client
}

type AuthenticationCookies struct {
	Cookies []*http.Cookie
}

func (app *App) login(username string, password string) {
	client := app.Client

	loginFormURL := baseURL + "/SiteManagerWeb/login.jsp"
	fmt.Printf("Acessando formulario de login.\n")
	resp, err := client.Get(loginFormURL)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("Formulario de login lido com sucesso.\n")

	data := url.Values{
		"j_username": {username},
		"j_password": {password},
	}

	fmt.Printf("Postando dados para o Formulario.\n")
	loginPostURL := baseURL + "/SiteManagerWeb/j_security_check"
	resp, err = client.PostForm(loginPostURL, data)
	if err != nil {
		// erro fatal, não consegui acessar a página
		log.Fatalln(err)
	}
	fmt.Printf("Dados postados com sucesso.\n")

	// Validar se o login deu certo
	if resp.StatusCode != http.StatusOK {
		log.Fatalln(errors.Wrap(errors.New("Erro "+strconv.Itoa(resp.StatusCode)+" na requisição"), "bla"))
	}
	//	pretty.Println(resp)

	defer resp.Body.Close()
	return
}

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
