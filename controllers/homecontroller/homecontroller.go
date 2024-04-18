package homecontroller

import (
	"html/template"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
