package main

import (
	_ "Bank/routers"

	"github.com/astaxie/beego"
	_ "github.com/beego/i18n"
)

func main() {
	beego.Run()
}
