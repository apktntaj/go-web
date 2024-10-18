package wiki

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)

	t, _ := template.ParseFiles("wiki/view.html")
	t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, _ := loadPage(title)

	t, _ := template.ParseFiles("wiki/edit.html")
	t.Execute(w, p)
}
