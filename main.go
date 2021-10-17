package main

import (
	"github.com/docker/docker/client"
	"github.com/gobuffalo/packr/v2"
	c "github.com/nickthetramp/portphosis/configurationManager"
	h "github.com/nickthetramp/portphosis/handler"
	"log"
	"net/http"
)

func main() {
	static := packr.New("static", "./static")
	configurations := packr.New("configurations", "./configurations")

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	manager := c.ConfigurationManager{
		Box:    configurations,
		Client: cli,
	}

	handler := h.Handler{
		Manager: manager,
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(static)))

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/config", handler.ConfigHandler)
	http.HandleFunc("/toggle", handler.ToggleContainerHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
