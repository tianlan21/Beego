package controllers

import (
	"github.com/astaxie/beego"
)

// MainController definition.
type RegisterController struct {
	beego.Controller
}

// Get method.
func (c *RegisterController) Get() {
	c.TplName = "register.html"
	c.Render()
}
