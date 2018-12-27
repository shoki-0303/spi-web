package controllers

import (
	"net/http"
	"spi-web/app/controllers/helpers"
	"spi-web/app/models"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// AdminMiddleWare : done before processeing related to admin
func AdminMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func AdminRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}

func AdminLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func AdminUpdate(c echo.Context) error {
	token := c.QueryParam("token")
	data := map[string]string{
		"token": token,
	}
	return c.Render(http.StatusOK, "update.html", data)
}

func AdminCreateUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	adminUser := &models.AdminUser{
		Name:           name,
		Email:          email,
		HashedPassword: string(hashedPassword),
	}

	err = models.CreateAdminUser(adminUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	//user is created and create token
	endpoint := "/admin/restricted"
	jwt := helpers.CreateJWTtoken(adminUser)
	url := helpers.CheckURL(endpoint, adminUser.Name, jwt)
	return c.Redirect(http.StatusSeeOther, url)
}

func ConfirmAdminUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	isComfirmed, err, adminUser := models.ConfirmAdminUser(email, password)
	if isComfirmed == true {
		//user is confirmed and create jwtToken
		endpoint := "/admin/restricted"
		jwt := helpers.CreateJWTtoken(&adminUser)
		url := helpers.CheckURL(endpoint, adminUser.Name, jwt)
		return c.Redirect(http.StatusSeeOther, url)
	}
	return echo.NewHTTPError(http.StatusUnauthorized, err)
}

func UpdateAdminUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	oldname := claims["name"].(string)

	rename := c.FormValue("name")
	adminUser, err := models.UpdateAdminUser(rename, oldname)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, nil)
	}

	//update is successed
	endpoint := "/admin/restricted"
	jwt := c.QueryParam("token")
	url := helpers.CheckURL(endpoint, adminUser.Name, jwt)
	return c.Redirect(http.StatusSeeOther, url)
}

func ShowAdminUser(c echo.Context) error {
	name := c.Param("name")
	adminUser, err := models.GetAdminUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	token := c.QueryParam("token")
	data := map[string]string{
		"adminUserName": adminUser.Name,
		"token":         token,
	}
	return c.Render(http.StatusOK, "adminUser.html", data)
}
