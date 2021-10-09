package main

import (
	"github.com/gobuffalo/packr/v2"
	c "github.com/nickthetramp/portphosis/configurationManager"
	h "github.com/nickthetramp/portphosis/handler"
	"log"
	"net/http"
)


func main() {
	static := packr.New("static","./static")
	configurations := packr.New("configurations","./configurations")

	manager := c.ConfigurationManager{ Box: configurations}
	handler := h.Handler{
		Manager: manager,
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(static)))
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/config", handler.ConfigHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
