package main

import (
	_ "gonganju/models"
	_ "gonganju/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

