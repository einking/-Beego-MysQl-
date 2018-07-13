package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type AdminController struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:newlife@tcp(127.0.0.1:3306)/Bank?charset=utf8", 30)
}

func (this *AdminController) Which() {
	Choose := this.GetString("Choose")
	beego.Debug(Choose)
	switch Choose {
	case "EstablishAccount":
		this.Redirect("/ad/establishaccounthtml", 302)
	case "CloseAccount":
		this.Redirect("/ad/closeaccounthtml", 302)
	case "RegisterCard":
		this.Redirect("/ad/registercardhtml", 302)
	case "CheckAccount":
		this.Redirect("/ad/checkaccounthtml", 302)
	}
}

func (this *AdminController) EstablishAccountHtml() {
	this.TplName = "establishaccounthtml.html"
}

func (this *AdminController) EstablishAccount() {
	UserId := this.GetString("UserId")
	Uname := this.GetString("Uname")
	Upassword := this.GetString("Upassword")

	//beego.Debug(UserId)
	ormm := orm.NewOrm()
	ormm.Using("default")

	res, _ := ormm.Raw("INSERT INTO User (UserId,Uname,Upassword,Confidence) VALUES(?,?,?,?)", UserId, Uname, Upassword, 5).Exec()
	beego.Debug(res)
	/*if err == nil {
		num, _ := res.RowsAffected()
		beego.Debug("mysql row affected nums: ", num)
		//this.Ctx.SetCookie("Success", "EstablishAccount Successful", 100, '/')
		//this.Redirect("/ad/adsuccess", 302)
	} else {
		beego.Debug(err)
		this.Ctx.SetCookie("Success", "error ", 100, '/')
		this.Redirect("/ad/adsuccess", 302)
		beego.Debug(UserId + Uname + Upassword)
	}*/
	var list []orm.ParamsList
	Result := ""
	num, err := ormm.Raw("SELECT * FROM User WHERE  UserId = ?", UserId).ValuesList(&list)
	if err == nil {
		fmt.Print(num)
		for i := range list {
			for j := range list[i] {
				Result += list[i][j].(string)
				//beego.Debug(list[i][j])
				Result += "--"
			}

		}
		beego.Debug(Result)
		this.Ctx.SetCookie("Success", Result, 100, '/')
		this.Redirect("/ad/adsuccess", 302)
	}
	//this.Redirect("/ad/registercardhtml", 302)
}

func (this *AdminController) RegisterCardHtml() {
	this.TplName = "registercardhtml.html"
}

func (this *AdminController) RegisterCard() {
	CardNum := this.GetString("CardNum")
	UserId := this.GetString("UserId")
	Cpassword := this.GetString("Cpassword")
	Brand := this.GetString("Brand")
	OverDraft := this.GetString("OverDraft")
	Remain := this.GetString("Remain")
	Loan := this.GetString("Loan")
	LoanBegin := this.GetString("LoanBegin")
	LoanTime := this.GetString("LoanTime")
	Rate := this.GetString("Rate")

	beego.Debug(UserId)
	ormm := orm.NewOrm()
	ormm.Using("default")
	//var list []orm.ParamsList
	res, _ := ormm.Raw("INSERT INTO Card (CardNum,UserId,Cpassword,Brand,OverDraft,Remain,Loan,LoanBegin,LoanTime,Rate) VALUES(?,?,?,?,?,?,?,?,?,?)",
		CardNum, UserId, Cpassword, Brand, OverDraft, Remain, Loan, LoanBegin, LoanTime, Rate).Exec()
	beego.Debug(res)
	/*if err == nil {
		num, _ := res.RowsAffected()
		this.Ctx.SetCookie("Success", "RegisterCard Successful", 100, '/')
		this.Redirect("/ad/adsuccess", 302)
		beego.Debug("mysql row affected nums: ", num)
	} else {
		this.Ctx.SetCookie("Success", "error ", 100, '/')
		this.Redirect("/ad/adsuccess", 302)
		beego.Debug(err)
		beego.Debug(UserId)

	}*/
	beego.Debug(res)
	var list []orm.ParamsList
	Result := ""
	num, err := ormm.Raw("SELECT * FROM Card WHERE  CardNum = ?", CardNum).ValuesList(&list)
	if err == nil {
		fmt.Print(num)
		for i := range list {
			for j := range list[i] {
				Result += list[i][j].(string)
				//beego.Debug(list[i][j])
				Result += "--"
			}

		}
		beego.Debug(Result)
		this.Ctx.SetCookie("Success", Result, 100, '/')
		this.Redirect("/ad/adsuccess", 302)
	}
	//this.Redirect("/ad/admin", 302)
}

func (this *AdminController) CloseAccountHtml() {
	this.TplName = "closeaccounthtml.html"
}

func (this *AdminController) CloseAccount() {
	UserId := this.GetString("UserId")
	CardNum := this.GetString("CardNum")

	ormm := orm.NewOrm()
	ormm.Using("default")

	var list orm.ParamsList
	res, err := ormm.Raw("SELECT * FROM (SELECT UserId,CardNum FROM User NATURAL JOIN Card) AS TEMP WHERE UserId=? AND CardNum = ?",
		UserId, CardNum).ValuesFlat(&list)
	beego.Debug(res)
	beego.Debug(list)
	if err == nil {
		resone, err := ormm.Raw("DELETE FROM User WHERE UserId = ?", UserId).Exec()
		if err == nil {
			beego.Debug(err)
			num, _ := resone.RowsAffected()
			beego.Debug("mysql row affected nums: ", num)

			this.Ctx.SetCookie("Success", "CloseAccount Successful", 100, '/')
			this.Redirect("/ad/adsuccess", 302)
		}
	} else {
		this.Ctx.SetCookie("Success", "error ", 100, '/')
		this.Redirect("/ad/adsuccess", 302)
		this.Redirect("/error", 302)
	}
	this.Redirect("/ad/adsuccess", 302)
	this.Redirect("/error", 302)
}

func (this *AdminController) CheckAccountHtml() {
	this.TplName = "checkaccounthtml.html"
}
func (this *AdminController) CheckAccount() {
	CardNum := this.GetString("CardNum")
	Name := this.GetString("Uname")
	Result := ""

	ormm := orm.NewOrm()
	var list []orm.ParamsList
	if CardNum != "" {
		res, err := ormm.Raw("SELECT * FROM History WHERE  CardNum = ?",
			CardNum).ValuesList(&list)
		if err == nil {
			fmt.Print(res)
			for i := range list {
				for j := range list[i] {
					Result += list[i][j].(string)
					//beego.Debug(list[i][j])
				
				}
				Result += "=="
			}
			beego.Debug(Result)
			this.Ctx.SetCookie("Success", Result, 100, '/')
			this.Redirect("/ad/adsuccess", 302)
		}
	}
	if Name != "" {
		res, err := ormm.Raw("SELECT * FROM User WHERE Uname LIKE '?%'", Name).ValuesList(&list)
		beego.Debug(err)
		if err == nil {
			for i := range list {
				fmt.Print(res)
				for j := range list[i] {
					Result += list[i][j].(string)
					//beego.Debug(list[i][j])
				}
				//Result += " "
			}
			this.Ctx.SetCookie("Success", Result, 100, '/')
			this.Redirect("/ad/adsuccess", 302)
		}
		this.Ctx.SetCookie("Success", "Find is not supported by Go", 100, '/')
		this.Redirect("/ad/adsuccess", 302)
	}
}

func (this *AdminController) LogIn() {
	this.Data["Website"] = "hello"
	this.TplName = "admin.html"
}

func (this *AdminController) Success() {
	this.Data["Success"] = this.Ctx.GetCookie("Success")
	this.TplName = "adsuccess.html"
}
