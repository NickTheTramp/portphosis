package handler

import (
	"html/template"
	"net/http"
)

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "Title", Configurations: h.Manager.GetConfigurations() }

	t, _ := template.ParseFiles("templates/index.tmpl.html", "templates/base.tmpl.html")
	t.Execute(w, p)
}

func (h *Handler) ConfigHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	if !params.Has("name") {
		w.Write([]byte("No name provided"))
	}
	name := params.Get("name")

	config, err := h.Manager.FindConfiguration(name)
	if err != nil {
		println([]byte(err.Error()))
		w.Write([]byte(err.Error()))
	}

	p := Page{Title: name, Configurations: h.Manager.GetConfigurations(), Configuration: config}

	t, _ := template.ParseFiles("templates/config.tmpl.html", "templates/base.tmpl.html")
	t.Execute(w, p)
}