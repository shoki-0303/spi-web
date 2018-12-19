package controllers

import (
	"fmt"
	"net/http"
	"spi-web/app/models"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

// AdminMiddleWare : done before processeing related to admin
func AdminMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	//"ここにadminUserを確かめる処理を書く"
	return next
}

func AdminRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
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
	redirectURL := fmt.Sprintf("/admin/%s", adminUser.Name)
	return c.Redirect(http.StatusSeeOther, redirectURL)
}

func ShowAdminUser(c echo.Context) error {
	name := c.Param("name")
	data := map[string]string{
		"adminUserName": name,
	}
	return c.Render(http.StatusOK, "adminUser.html", data)
}
