package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Yefhem/student-syllabus/helpers"
	"github.com/Yefhem/student-syllabus/model"
	"github.com/Yefhem/student-syllabus/service"
	"github.com/labstack/echo/v4"
)

type AuthController interface {
	LoginPage(c echo.Context) error
	RegisterPage(c echo.Context) error

	CreateAccount(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

type authController struct {
	userService    service.UserService
	sessionService service.SessionService
	alertService   service.AlertService
}

func NewAuthController(userService service.UserService, sessionService service.SessionService, alertService service.AlertService) AuthController {
	return &authController{
		userService:    userService,
		sessionService: sessionService,
		alertService:   alertService,
	}
}

var authMap = make(map[string]interface{})

// --------------------> Pages

// ----------> Login Page
func (cont *authController) LoginPage(c echo.Context) error {
	results, err := helpers.Include("userops/login")
	if err != nil {
		log.Fatal(err)
	}
	view, err := template.ParseFiles(results...)
	if err != nil {
		fmt.Println(err)
	}

	authMap["Alert"] = cont.alertService.GetAlert(c)

	return view.ExecuteTemplate(c.Response(), "index", authMap)
}

// ----------> Register Page
func (cont *authController) RegisterPage(c echo.Context) error {
	results, err := helpers.Include("userops/register")
	if err != nil {
		log.Fatal(err)
	}
	view, err := template.ParseFiles(results...)
	if err != nil {
		fmt.Println(err)
	}

	authMap["Alert"] = cont.alertService.GetAlert(c)

	return view.ExecuteTemplate(c.Response(), "index", authMap)
}

// --------------------> Methods

// ----------> Create Account
func (cont *authController) CreateAccount(c echo.Context) error {

	name := c.FormValue("name")
	email := c.FormValue("email")

	if !helpers.IsEmailValid(email) {
		cont.alertService.SetAlert(c, "Email Hatalı!")
		return c.Redirect(http.StatusSeeOther, "/register")
	}

	// ---- Email var mı yok mu sorgusu at

	password := c.FormValue("pass")
	repeatPassword := c.FormValue("rep-pass")

	if password != repeatPassword {
		cont.alertService.SetAlert(c, "Passwordlar eşleşmiyor")
		return c.Redirect(http.StatusSeeOther, "/register")
	}

	userDTO := model.UserDTO{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := cont.userService.CreateUser(userDTO); err != nil {
		log.Println(err)
		cont.alertService.SetAlert(c, "Kayıt Esnasında Bir Hata Oluştu!")
		return c.Redirect(http.StatusSeeOther, "/register")
	}

	cont.alertService.SetAlert(c, "Kayıt Başarılı!")
	return c.Redirect(http.StatusSeeOther, "/login")
}

// ----------> Login

func (cont *authController) Login(c echo.Context) error {

	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := cont.userService.FindUserByEmail(email)
	if err != nil {
		log.Println(err)
		cont.alertService.SetAlert(c, "Yanlış Kullanıcı Adı veya Şifre!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	storedUserPass := user.Password

	if cont.userService.VerifyPassword(storedUserPass, password) {
		// --- login
		log.Println("login girdi")
		cont.sessionService.SetUser(c, user.Name, user.Email, user.Password)
		cont.alertService.SetAlert(c, "Welcome")
		return c.Redirect(http.StatusSeeOther, "/tasks")
	} else {
		// --- denied
		cont.alertService.SetAlert(c, "Yanlış Kullanıcı Adı veya Şifre!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

}

// ----------> Logout
func (cont *authController) Logout(c echo.Context) error {
	cont.sessionService.RemoveUser(c)
	cont.alertService.SetAlert(c, "Bye Bye..")
	return c.Redirect(http.StatusSeeOther, "/login")
}
