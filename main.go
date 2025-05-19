package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var tpl *template.Template

func main() {
    tpl = template.Must(template.ParseGlob("templates/*.html"))
    if tpl == nil {
        panic("no tpl???")
    }
	
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", indexHandler)

	err := http.ListenAndServeTLS("127.0.0.1:3000", "certs/domain.cert.pem", "certs/private.key.pem", nil)
	if err != nil {
		panic("could not start server:" + err.Error())
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://cdn.discordapp.com/emojis/1255615186346184796.webp", http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path[len("/"):] != "" {
        http.ServeFile(w, r, filepath.Join("static/", r.URL.Path))
        return
    }

    err := tpl.ExecuteTemplate(w, "index.html", nil)
    if err != nil {
        http.Error(w, "Error Rendering Template", http.StatusInternalServerError)
    }
}
