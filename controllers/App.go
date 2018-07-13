package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type AppController struct {
	beego.Controller
}

type User struct {
	UserId     int
	Uname      string
	Upassword  string
	Confidence int
}

type Admin struct {
	AdminId   string
	Aname     string
	Apassword string
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:newlife@tcp(127.0.0.1:3306)/Bank?charset=utf8", 30)
}

func (this *AppController) Get() {
	this.TplName = "index.html"
}

func (this *AppController) LogIn() {
	ormm := orm.NewOrm()
	ormm.Using("default")

	usernumber := this.GetString("usernumber")
	password := this.GetString("password")
	role := this.GetString("role")
	beego.Debug(usernumber + role + password)
	fmt.Println(role + password)
	if len(usernumber) == 0 {
		this.Data["Website"] = "Wrong"
		this.TplName = "error.html"
	}
	switch role {
	case "User":
		var list orm.ParamsList
		num, err := ormm.Raw("SELECT * FROM User  WHERE UserId = ? AND Upassword=?", usernumber, password).ValuesFlat(&list)
		beego.Debug(err)
		beego.Debug(list)
		if err == nil && num > 0 {
			this.Redirect("/us/login", 302)

		} else {
			this.Ctx.SetCookie("Error", "Wrong With Log In", 100, '/')
			this.Redirect("/error", 302)
		}

	case "Admin":
		var list orm.ParamsList
		num, err := ormm.Raw("SELECT * FROM Admin WHERE AdminId = ? AND Apassword=?", usernumber, password).ValuesFlat(&list)
		beego.Debug(err)
		beego.Debug(list)
		if err == nil && num > 0 {
			this.Redirect("/ad/login", 302)

		} else {
			this.Ctx.SetCookie("Error", "Wrong With Log In", 100, '/')
			this.Redirect("/error", 302)
		}
	}
}

func (this *AppController) Error() {
	this.Data["Error"] = this.Ctx.GetCookie("Error")
	this.TplName = "error.html"
}

func (this *AppController) Success() {
	this.Data["Success"] = this.Ctx.GetCookie("Success")
	beego.Debug("this.Ctx.GetCookie")
	beego.Debug(this.Ctx.GetCookie("Success"))
	this.TplName = "success.html"
}
