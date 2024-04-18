package divisioncontroller

import (
	"html/template"
	"indodev-test/entities"
	"indodev-test/models/divisionsmodel"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	divisions := divisionsmodel.GetAll()
	data := map[string]any{
		"divisions": divisions,
	}

	temp, err := template.ParseFiles("views/divisions/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/divisions/addform.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var division entities.Divisions

		division.DivisionId = r.FormValue("divisionid")
		division.DivisionName = r.FormValue("divisionname")
		division.DepartmentId = r.FormValue("departmentid")

		if ok := divisionsmodel.Insert(division); !ok {
			temp, _ := template.ParseFiles("views/divisions/addform.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/divisions", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/divisions/editform.html")
		if err != nil {
			panic(err)
		}

		id := r.URL.Query().Get("id")

		division := divisionsmodel.Detail(id)
		data := map[string]any{
			"division": division,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var division entities.Divisions

		id := r.FormValue("id")
		division.DivisionName = r.FormValue("name")

		if ok := divisionsmodel.Update(id, division); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/divisions", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := divisionsmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/divisions", http.StatusSeeOther)
}
