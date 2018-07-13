package routers

import (
	"Bank/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.AppController{})
	beego.Router("/login", &controllers.AppController{}, "post:LogIn")
	beego.Router("/error", &controllers.AppController{}, "get:Error")
	beego.Router("/success", &controllers.AppController{}, "get:Success")

	beego.Router("/ad", &controllers.AdminController{})
	beego.Router("/ad/login", &controllers.AdminController{}, "get:LogIn")
	beego.Router("/ad/adsuccess", &controllers.AdminController{}, "get:Success")
	beego.Router("/ad/which", &controllers.AdminController{}, "post:Which")
	beego.Router("/ad/establishaccounthtml", &controllers.AdminController{}, "get:EstablishAccountHtml")
	beego.Router("/ad/establishaccount", &controllers.AdminController{}, "post:EstablishAccount")
	beego.Router("/ad/registercardhtml", &controllers.AdminController{}, "get:RegisterCardHtml")
	beego.Router("/ad/registercard", &controllers.AdminController{}, "post:RegisterCard")
	beego.Router("/ad/closeaccounthtml", &controllers.AdminController{}, "get:CloseAccountHtml")
	beego.Router("/ad/closeaccount", &controllers.AdminController{}, "post:CloseAccount")
	beego.Router("/ad/checkaccounthtml", &controllers.AdminController{}, "get:CheckAccountHtml")
	beego.Router("/ad/checkaccount", &controllers.AdminController{}, "post:CheckAccount")

	beego.Router("/us", &controllers.UserController{})
	beego.Router("/us/login", &controllers.UserController{}, "get:LogIn")
	beego.Router("/us/which", &controllers.UserController{}, "post:Which")

	beego.Router("/us/deposithtml", &controllers.UserController{}, "get:DepositHtml")
	beego.Router("/us/deposit", &controllers.UserController{}, "post:Deposit")
	beego.Router("/us/usdeposithtml", &controllers.UserController{}, "get:UsDepositHtml")
	beego.Router("/us/usdeposit", &controllers.UserController{}, "post:UsDeposit")
	beego.Router("/us/withdralshtml", &controllers.UserController{}, "get:WithdralsHtml")
	beego.Router("/us/withdrals", &controllers.UserController{}, "post:Withdrals")
	beego.Router("/us/uswithdralshtml", &controllers.UserController{}, "get:UsWithdralsHtml")
	beego.Router("/us/uswithdrals", &controllers.UserController{}, "post:UsWithdrals")
	beego.Router("/us/loanhtml", &controllers.UserController{}, "get:LoanHtml")
	beego.Router("/us/loan", &controllers.UserController{}, "post:Loan")
	beego.Router("/us/usloanhtml", &controllers.UserController{}, "get:UsLoanHtml")
	beego.Router("/us/usloan", &controllers.UserController{}, "post:UsLoan")
	beego.Router("/us/transferhtml", &controllers.UserController{}, "get:TransferHtml")
	beego.Router("/us/transfer", &controllers.UserController{}, "post:Transfer")
	beego.Router("/us/ustransferhtml", &controllers.UserController{}, "get:UsTransferHtml")
	beego.Router("/us/ustransfer", &controllers.UserController{}, "post:UsTransfer")
	beego.Router("/us/repaymenthtml", &controllers.UserController{}, "get:RepaymentHtml")
	beego.Router("/us/repayment", &controllers.UserController{}, "post:Repayment")
	beego.Router("/us/usrepaymenthtml", &controllers.UserController{}, "get:UsRepaymentHtml")
	beego.Router("/us/usrepayment", &controllers.UserController{}, "post:UsRepayment")
}
