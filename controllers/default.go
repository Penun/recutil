package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type BaseResponse struct {
	Success bool `json:"success"`
	Error string `json:"error_code"`
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}
