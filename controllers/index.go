package controllers

import (
	"github.com/astaxie/beego"
)

type Index struct {
	beego.Controller
}

func (c *Index) Get() {
	c.TplNames = "index.html"
}

