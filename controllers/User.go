package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:newlife@tcp(127.0.0.1:3306)/Bank?charset=utf8", 30)
}

type UserController struct {
	beego.Controller
}

func (this *UserController) Which() {
	Choose := this.GetString("Choose")
	beego.Debug(Choose)
	switch Choose {
	case "Deposit":
		this.Redirect("/us/deposithtml", 302)
	case "Withdrals":
		this.Redirect("/us/withdralshtml", 302)
	case "Loan":
		this.Redirect("/us/loanhtml", 302)
	case "Transfer":
		this.Redirect("/us/transferhtml", 302)
	case "Repayment":
		this.Redirect("/us/repaymenthtml", 302)
	}
}

func (this *UserController) DepositHtml() {

	this.TplName = "deposithtml.html"
}
func (this *UserController) Deposit() {
	Brand := this.GetString("Brand")
	CardNum := this.GetString("CardNum")
	Cpassword := this.GetString("Upassword")
	beego.Debug(Cpassword)
	this.Ctx.SetCookie("CardNum", CardNum, 100, "/")

	ormm := orm.NewOrm()
	ormm.Using("default")

	/*var list orm.ParamsList
	res, err := ormm.Raw("SELECT * FROM Card WHERE CardNum=? AND Cpassword=?string Brand=?", CardNum, Cpassword, Brand).ValuesFlat(&list)*/
	res, err := ormm.Raw("SELECT * FROM Card WHERE CardNum=? AND Cpassword=? AND Brand=?", CardNum, Cpassword, Brand).Exec()
	if err != nil {
		beego.Debug(Cpassword)
		this.Ctx.WriteString("Sorry,your input has error")
	} else {
		beego.Debug(err)
		beego.Debug(res)
		this.Redirect("/us/usdeposithtml", 302)
	}
}
func (this *UserController) UsDepositHtml() {
	this.TplName = "usdeposithtml.html"
}
func (this *UserController) UsDeposit() {
	Money := this.GetString("Money")
	CardNum := this.Ctx.GetCookie("CardNum")

	ormm := orm.NewOrm()
	ormm.Using("default")

	var list orm.ParamsList
	num, err := ormm.Raw("SELECT Remain FROM Card WHERE CardNum =?", CardNum).ValuesFlat(&list)
	beego.Debug(list[0])
	beego.Debug(num)
	if err == nil && num > 0 {
		RemainInt, err := strconv.Atoi(list[0].(string))
		if err != nil {
			this.Redirect("/error", 302)
		}
		MoneyInt, err := strconv.Atoi(Money)
		if err != nil {
			this.Redirect("/error", 302)
		}
		MoneyInt = MoneyInt + RemainInt
		Money = strconv.Itoa(MoneyInt)
	}
	ress, err := ormm.Raw("UPDATE Card SET Remain=? WHERE CardNum =?;", Money, CardNum).Exec()
	resss, err := ormm.Raw("INSERT INTO History (CardNum,Operate,Timee) VALUES(?,?,?)", CardNum, "Deposit", time.Now()).Exec()
	if err == nil {
		beego.Debug(ress)
		beego.Debug(resss)
		this.Redirect("/us/login", 302)
	} else {
		this.Redirect("/error", 302)
	}
	var listtwo []orm.ParamsList
	Result := ""
	num, err = ormm.Raw("SELECT * FROM Card WHERE  CardNum = ?", CardNum).ValuesList(&listtwo)
	if err == nil {
		fmt.Print(num)
		for i := range listtwo {
			for j := range listtwo[i] {
				Result += listtwo[i][j].(string)
				//beego.Debug(list[i][j])
				Result += "--"
			}

		}
		beego.Debug(Result)
		this.Ctx.SetCookie("Success", Result, 100, '/')
		this.Redirect("/ad/adsuccess", 302)
	}
}

func (this *UserController) WithdralsHtml() {
	this.TplName = "withdralshtml.html"
}
func (this *UserController) Withdrals() {
	Brand := this.GetString("Brand")
	CardNum := this.GetString("CardNum")
	Cpassword := this.GetString("Password")

	this.Ctx.SetCookie("CardNum", CardNum, 100, "/")

	ormm := orm.NewOrm()
	ormm.Using("default")

	res, err := ormm.Raw("SELECT * FROM Card WHERE CardNum=? AND Cpassword=? AND Brand=?", CardNum, Cpassword, Brand).Exec()
	if err == nil {
		var list orm.ParamsList
		var listone orm.ParamsList
		num, err := ormm.Raw("SELECT Remain, OverDraft FROM Card WHERE CardNum = ?", CardNum).ValuesFlat(&list)
		//beego.Debug(list)
		fmt.Print(list)
		//beego.Debug(num)
		fmt.Print(num)
		//beego.Debug("ree")
		//beego.Debug(err)
		fmt.Print(err)
		if list[1].(string) != "0" {
			beego.Debug("Confidence")
			beego.Debug(list)
			//this.Ctx.WriteString("Can not OverGraft and Remain is" + list[0].(string))
			this.Ctx.SetCookie("OverDraft", "Cannot", 100, "/")
			this.Ctx.SetCookie("Remain", list[0].(string), 100, "/")
			this.Redirect("/us/uswithdralshtml", 302)
		} else {
			num, err := ormm.Raw("SELECT Confidence FROM (SELECT * FROM User NATURAL JOIN Card) AS TEMP WHERE CardNum = ?",
				CardNum).ValuesFlat(&listone)
			beego.Debug(listone)
			Confidence, err := strconv.Atoi(listone[0].(string))
			if err == nil {
				fmt.Print(num)
			}
			//this.Ctx.WriteString("Can OverGraft is " + strconv.Itoa(Confidence*1000) + " and Remain is " + list[0].(string))
			//this.Ctx.WriteString("Can not OverGraft and Remain is" + list[0].(string))
			
			
			this.Ctx.SetCookie("OverDraft", strconv.Itoa(Confidence*1000), 100, "/")
			this.Ctx.SetCookie("Remain", list[0].(string), 100, "/")
			this.Redirect("/us/uswithdralshtml", 302)
		}
		beego.Debug(res)
		//this.Redirect("/us/withdralshtml", 302)
	} else {
		this.Redirect("/error", 302)
	}
}
func (this *UserController) UsWithdralsHtml() {
	this.Data["OverDraft"] = this.Ctx.GetCookie("OverDraft")
	this.Data["Remain"] = this.Ctx.GetCookie("Remain")
	this.TplName = "uswithdralshtml.html"
}
func (this *UserController) UsWithdrals() {
	Money := this.GetString("Money")
	CardNum := this.Ctx.GetCookie("CardNum")

	ormm := orm.NewOrm()
	ormm.Using("default")

	var list orm.ParamsList
	num, err := ormm.Raw("SELECT Remain FROM Card WHERE CardNum =?", CardNum).ValuesFlat(&list)
	beego.Debug(list[0])
	beego.Debug(num)
	var MoneyInt int
	if err == nil && num > 0 {
		RemainInt, err := strconv.Atoi(list[0].(string))
		if err != nil {
			this.Redirect("/error", 302)
		}
		MoneyInt, err := strconv.Atoi(Money)
		if err != nil {
			this.Redirect("/error", 302)
		}
		MoneyInt = RemainInt - MoneyInt
		Money = strconv.Itoa(MoneyInt)
	}
	if MoneyInt < 0{
		res, _ := ormm.Raw("UPDATE Card SET OverDraft=?,Remain=? WHERE CardNum =?;", "1",Money, CardNum).Exec()
		beego.Debug(res)
	}else{
		res, _ := ormm.Raw("UPDATE Card SET Remain=? WHERE CardNum =?;", Money, CardNum).Exec()
		beego.Debug(res)
	}
	res, err := ormm.Raw("INSERT INTO History (CardNum,Operate,Timee) VALUES(?,?,?)", CardNum, "Withdrals", time.Now()).Exec()
	/*if err == nil {
		beego.Debug(res)
		this.Redirect("/us/login", 302)
	} else {
		this.Redirect("/error", 302)
	}*/
	beego.Debug(res)
	var listtwo []orm.ParamsList
	Result := ""
	num, err = ormm.Raw("SELECT * FROM Card WHERE  CardNum = ?", CardNum).ValuesList(&listtwo)
	if err == nil {
		fmt.Print(num)
		for i := range listtwo {
			for j := range listtwo[i] {
				Result += listtwo[i][j].(string)
				//beego.Debug(list[i][j])
				Result += "--"
			}

		}
		beego.Debug(Result)
		this.Ctx.SetCookie("Success", Result, 100, '/')
		this.Redirect("/ad/adsuccess", 302)
	}
}

func (this *UserController) LoanHtml() {
	this.TplName = "loanhtml.html"
}
func (this *UserController) Loan() {
	Brand := this.GetString("Brand")
	CardNum := this.GetString("CardNum")
	Cpassword := this.GetString("Password")

	this.Ctx.SetCookie("CardNum", CardNum, 100, "/")

	ormm := orm.NewOrm()
	ormm.Using("default")

	res, err := ormm.Raw("SELECT * FROM Card WHERE CardNum=? AND Cpassword=? AND Brand=?", CardNum, Cpassword, Brand).Exec()
	if err == nil {
		var list orm.ParamsList
		var listone orm.ParamsList
		num, err := ormm.Raw("SELECT Loan FROM Card WHERE CardNum = ?", CardNum).ValuesFlat(&list)
		beego.Debug(list)
		fmt.Print(list)
		//beego.Debug(num)
		fmt.Print(num)
		//beego.Debug("ree")
		//beego.Debug(err)
		fmt.Print(err)
		if list[0].(string) == "NO" {
			num, err := ormm.Raw("SELECT Confidence FROM (SELECT * FROM User NATURAL JOIN Card) AS TEMP WHERE CardNum = ?",
				CardNum).ValuesFlat(&listone)
			beego.Debug(listone)
			Confidence, err := strconv.Atoi(listone[0].(string))
			if err == nil {
				fmt.Print(num)
			}
			//this.Ctx.WriteString("Can OverGraft is " + strconv.Itoa(Confidence*1000) + " and Remain is " + list[0].(string))
			//this.Ctx.WriteString("Can not OverGraft and Remain is" + list[0].(string))
			this.Ctx.SetCookie("Loan", strconv.Itoa(Confidence*10000), 100, "/")
			fmt.Print(Confidence)
			this.Redirect("/us/usloanhtml", 302)
		} else {
			this.Ctx.WriteString(" This Card is Loaned,cannot operator again")
		}
		//beego.Debug(res)
		fmt.Print(res)
	} else {
		this.Redirect("/error", 302)
	}
}
func (this *UserController) UsLoanHtml() {
	this.Data["Loan"] = this.Ctx.GetCookie("Loan")
	this.TplName = "usloanhtml.html"
}
func (this *UserController) UsLoan() {
	LoanTime := this.GetString("LoanTime")
	Rate := this.GetString("Rate")
	Money := this.GetString("Money")
	CardNum := this.Ctx.GetCookie("CardNum")

	var listone orm.ParamsList
	ormm := orm.NewOrm()

	beego.Debug(Rate)
	num, err := ormm.Raw("SELECT Remain FROM  Card WHERE CardNum = ?",
		CardNum).ValuesFlat(&listone)
	if err == nil {
		beego.Debug(num)
	}
	Remain,_:=strconv.Atoi(listone[0].(string))
	Loan_int,_:=strconv.Atoi(Money)
	Remain+=Loan_int
	Money_str:=strconv.Itoa(Remain)
	beego.Debug(listone)
	/*Year := time.Now().Year()
	Month := time.Now().Month()
	Day := time.Now().Day()*/
	res, _ := ormm.Raw("UPDATE Card SET Remain=?,Loan=? , LoanBegin=?,LoanTime=?,Rate=? WHERE CardNum = ?;",
		Money_str,Money, time.Now(), LoanTime, Rate, CardNum).Exec()
	res, _ = ormm.Raw("INSERT INTO History (CardNum,Operate,Timee) VALUES(?,?,?)", CardNum, "Loan", time.Now()).Exec()
	/*if err == nil {
		beego.Debug(res)
	}*/
	beego.Debug(res)
	var listtwo []orm.ParamsList
	Result := ""
	num, err = ormm.Raw("SELECT * FROM Card WHERE  CardNum = ?", CardNum).ValuesList(&listtwo)
	if err == nil {
		fmt.Print(num)
		for i := range listtwo {
			for j := range listtwo[i] {
				Result += listtwo[i][j].(string)
				//beego.Debug(list[i][j])
				Result += "--"
			}
		}
		beego.Debug(Result)
		this.Ctx.WriteString(Result)
	}
	this.Redirect("/us/login", 302)
}

func (this *UserController) TransferHtml() {
	this.TplName = "transferhtml.html"
}
func (this *UserController) Transfer() {
	Brand := this.GetString("Brand")
	CardNum := this.GetString("CardNum")
	Cpassword := this.GetString("Password")

	this.Ctx.SetCookie("CardNum", CardNum, 100, "/")

	ormm := orm.NewOrm()
	ormm.Using("default")

	res, err := ormm.Raw("SELECT * FROM Card WHERE CardNum=? AND Cpassword=? AND Brand = ?", CardNum, Cpassword, Brand).Exec()
	if err == nil {
		var list orm.ParamsList

		num, err := ormm.Raw("SELECT Remain FROM Card WHERE CardNum = ?", CardNum).ValuesFlat(&list)
		this.Ctx.SetCookie("Remain", list[0].(string), 100, "/")
		this.Ctx.SetCookie("CardNum", CardNum, 100, "/")
		this.Ctx.SetCookie("Brand", Brand, 100, "/")
		beego.Debug(list)
		fmt.Print(list)
		//beego.Debug(num)
		fmt.Print(num)
		//beego.Debug("ree")
		//beego.Debug(err)
		fmt.Print(err)
		this.Redirect("/us/ustransferhtml", 302)

		//beego.Debug(res)
		fmt.Print(res)
	} else {
		this.Redirect("/error", 302)
	}
}
func (this *UserController) UsTransferHtml() {
	this.Data["Remain"] = this.Ctx.GetCookie("Remain")
	this.TplName = "ustransferhtml.html"
}
func (this *UserController) UsTransfer() {
	CardNum := this.Ctx.GetCookie("CardNum")
	TranNum := this.GetString("TransferNumber")
	Money := this.GetString("Money")
	Brand := this.GetString("Brand")

	fmt.Print(Brand)

	ormm := orm.NewOrm()
	var list orm.ParamsList
	var listone orm.ParamsList
	num, err := ormm.Raw("SELECT Remain,Brand FROM Card WHERE CardNum = ?", CardNum).ValuesFlat(&list)
	res, err := ormm.Raw("INSERT INTO History (CardNum,Operate,Timee) VALUES(?,?,?)", CardNum, "Transfer", time.Now()).Exec()
	if err == nil {
		fmt.Print(num)
		fmt.Print(res)
	}
	fmt.Print(list)
	Brand_old := list[1].(string)
	Remain_old_int, _ := strconv.ParseFloat(list[0].(string), 16)
	Money_flo, _ := strconv.ParseFloat(Money, 32)
	//Money_int = float(Money_int)
	num, err = ormm.Raw("SELECT Remain, Brand FROM Card WHERE CardNum = ?", TranNum).ValuesFlat(&listone)
	Brand_new := list[1].(string)
	if Brand_old != Brand_new {
		Money_flo = Money_flo + Money_flo*0.05
		temp_old_int := Remain_old_int - Money_flo
		if temp_old_int < 0 {
			this.Ctx.SetCookie("Error", "The Remain is not enough", 100, "/")
			this.Redirect("/error", 302)
		}
		Money_old_str := strconv.FormatFloat(temp_old_int, 'E', -1, 64)

		res, err := ormm.Raw("UPDATE Card SET Remain=? WHERE CardNum =?;", Money_old_str, CardNum).Exec()

		Remain_new, _ := strconv.ParseFloat(listone[0].(string), 16)
		temp_old := Remain_new + Money_flo
		Money_new_str := strconv.FormatFloat(temp_old, 'E', -1, 64)

		res, err = ormm.Raw("UPDATE Card SET Remain=? WHERE CardNum =?;", Money_new_str, TranNum).Exec()
		if err == nil {
			fmt.Print(res)
		}
		this.Redirect("/us/login", 302)
	} else {
		temp_old_int := Remain_old_int - Money_flo
		Money_old_str := strconv.FormatFloat(temp_old_int, 'E', -1, 64)

		if temp_old_int < 0 {
			this.Ctx.SetCookie("Error", "The Remain is not enough", 100, "/")
			this.Redirect("/error", 302)
		}

		res, err := ormm.Raw("UPDATE Card SET Remain=? WHERE CardNum =?;", Money_old_str, CardNum).Exec()

		Remain_new, _ := strconv.ParseFloat(listone[0].(string), 16)
		temp_old := Remain_new + Money_flo
		Money_new_str := strconv.FormatFloat(temp_old, 'E', -1, 64)

		res, err = ormm.Raw("UPDATE Card SET Remain=? WHERE CardNum =?;", Money_new_str, TranNum).Exec()
		if err == nil {
			fmt.Print(res)
		}
		this.Ctx.SetCookie("Success", "Operate Successful", 100, "/")
		this.Redirect("/us/login", 302)
	}
	
}

func (this *UserController) RepaymentHtml() {
	this.TplName = "repaymenthtml.html"
}
func (this *UserController) Repayment() {
	Brand := this.GetString("Brand")
	CardNum := this.GetString("CardNum")
	Cpassword := this.GetString("Password")

	this.Ctx.SetCookie("CardNum", CardNum, 100, "/")

	ormm := orm.NewOrm()
	ormm.Using("default")

	res, err := ormm.Raw("SELECT * FROM Card WHERE CardNum=? AND Cpassword=? AND Brand = ?", CardNum, Cpassword, Brand).Exec()
	if err == nil {
		var list orm.ParamsList
		var listone orm.ParamsList
		num, err := ormm.Raw("SELECT Loan,Rate,LoanBegin,LoanTime FROM Card WHERE CardNum = ?", CardNum).ValuesFlat(&list)

		Rate_flo, _ := strconv.ParseFloat(list[1].(string), 16)
		Loan_flo, _ := strconv.ParseFloat(list[0].(string), 16)
		LoanBegin, _ := time.Parse(list[2].(string), "2006-01-02 15:04:05")
		LoanTime := list[3].(string)
		LoanTime_int, _ := strconv.Atoi(LoanTime)
		DuringYear := time.Now().Sub(LoanBegin)

		beego.Debug("DuringYear")
		beego.Debug(DuringYear)
		//fmt.Printf("%v year\n", )
		var year_int int
		year_temp := DuringYear.Hours() / 24 * 365
		if year_temp < 0 {
			year_int = 1
		} else {
			year_int = int(year_temp)
		}
		if year_int > LoanTime_int {
			num, err := ormm.Raw("SELECT UserId,Confidence FROM (SELECT * FROM User NATURAL JOIN Card) AS TEMP WHERE CardNum = ?",
				CardNum).ValuesFlat(&listone)
			beego.Debug(listone)
			UserId := listone[0].(string)
			Confidence, err := strconv.Atoi(listone[1].(string))
			Confidence = Confidence - 1
			res, err := ormm.Raw("UPDATE User SET Confidence=? WHERE UserId =?", Confidence, UserId).Exec()
			if err == nil {
				fmt.Print(num)
				fmt.Print(res)
			}
		}
		year_float := float64(year_int)
		Loan_flo = Loan_flo + Loan_flo*Rate_flo*year_float

		Loan := strconv.FormatFloat(Loan_flo, 'E', -1, 64)

		var (
			old = Loan
			new float64
		)
		n, err := fmt.Sscanf(old, "%e", &new)
		if err != nil {
			fmt.Println(err.Error())
		} else if 1 != n {
			fmt.Println("n is not one")
		}
		temp, err := strconv.Atoi(list[0].(string))
		n = n + temp

		fmt.Println(uint64(new))

		this.Ctx.SetCookie("Loan", strconv.Itoa(n), 100, "/")
		this.Ctx.SetCookie("CardNum", CardNum, 100, "/")
		this.Ctx.SetCookie("Brand", Brand, 100, "/")
		beego.Debug(list)
		fmt.Print(list)
		//beego.Debug(num)
		fmt.Print(num)
		//beego.Debug("ree")
		//beego.Debug(err)
		fmt.Print(err)
		this.Redirect("/us/usrepaymenthtml", 302)

		//beego.Debug(res)
		fmt.Print(res)
	} else {
		this.Redirect("/error", 302)
	}
}
func (this *UserController) UsRepaymentHtml() {
	this.Data["Loan"] = this.Ctx.GetCookie("Loan")
	this.TplName = "usrepaymenthtml.html"
}
func (this *UserController) UsRepayment() {
	Money := this.GetString("Money")
	CardNum := this.Ctx.GetCookie("CardNum")

	beego.Debug(Money)
	/*Year := time.Now().Year()
	Month := time.Now().Month()
	Day := time.Now().Day()*/
	ormm := orm.NewOrm()
	res, err := ormm.Raw("UPDATE Card SET Loan=? , LoanBegin=?,LoanTime=?,Rate=? WHERE CardNum = ?;",
		"NO", "0", "0", "0", CardNum).Exec()
	res, err = ormm.Raw("INSERT INTO History (CardNum,Operate,Timee) VALUES(?,?,?)", CardNum, "Repayment", time.Now()).Exec()

	if err == nil {
		beego.Debug(res)
	}
	beego.Debug(res)
	var listtwo []orm.ParamsList
	Result := ""
	num, err := ormm.Raw("SELECT * FROM Card WHERE  CardNum = ?", CardNum).ValuesList(&listtwo)
	if err == nil {
		fmt.Print(num)
		for i := range listtwo {
			for j := range listtwo[i] {
				Result += listtwo[i][j].(string)
				//beego.Debug(list[i][j])
				Result += "--"
			}
		}
		//beego.Debug(Result)
		//this.Ctx.SetCookie("Success", Result, 100, '/')
		//this.Redirect("/ad/adsuccess", 302)
		this.Ctx.WriteString("Congratulations ! Repayment Successful.")
	}
	this.Redirect("/us/login", 302)
}

func (this *UserController) LogIn() {
	this.Data["Success"] = this.Ctx.GetCookie("Success")
	this.TplName = "user.html"
}
