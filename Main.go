package main

import "net/http"
import (
	"flag"
	"time"
	"github.com/fatih/color"
	"html/template"
)

type parsetohtml struct {
	Data_1 string
	Data_2 string
}

var tmpl = template.Must(template.ParseGlob("template/*"))

func index(w http.ResponseWriter, r *http.Request) {
	var par = parsetohtml{}
	par.Data_1 = "coba 1"
	par.Data_2 = "coba 2"
	tmpl.ExecuteTemplate(w, "index", par)
}

func contact(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "contact", "")
}

func home(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "home", "")
}

func doAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "5" && password == "5" {
			http.Redirect(w, r, "result", 301)
		} else {
			http.Redirect(w, r, "err", 301)
		}

	}
}

func result(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "result", "")
}

func err(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "err", "")
}

func route() {
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/doAction", doAction)
	http.HandleFunc("/result", result)
	http.HandleFunc("/err", err)
}

func main() {
	route()
	var port = flag.String("port", "7033", "isi port")
	flag.Parse()
	var timenow = time.Now()
	c := color.New(color.FgGreen).Add(color.Underline)
	c.Printf("%s %s %s", timenow, "Berjalan di port", *port)
	http.ListenAndServe(":" + *port, nil)
}
