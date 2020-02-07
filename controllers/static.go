package controllers

import (
	"github.com/RudePupp/whopays/views"
)

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
	}
}

type Static struct {
	Home    *views.View
	Contact *views.View
}