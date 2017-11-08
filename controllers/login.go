package controllers

import (
	"github.com/astaxie/beego"
)

// MainController definition.
type LoginController struct {
	beego.Controller
}

// Get method.
func (c *LoginController) Get() {
	c.TplName = "login.html"
	c.Render()
}
