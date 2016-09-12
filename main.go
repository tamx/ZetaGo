package main

import (
	"controllers"
	"github.com/astaxie/beegae"
	_ "github.com/astaxie/beegae/session/appengine"
)

func init() {
	// beegae.TemplateLeft = "{{{"
	// beegae.TemplateRight = "}}}"
	// beegae.SessionOn = true
	beegae.Router("/", &controllers.BoardController{})
	beegae.Run()
}
