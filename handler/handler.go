package handler

import (
	"github.com/gobuffalo/packr/v2"
	c "github.com/nickthetramp/portphosis/configurationManager"
)

type Handler struct {
	Files packr.Box
	Manager c.ConfigurationManager
}

type Page struct {
	Title string
	Configurations []c.Configuration
	Configuration c.Configuration
}