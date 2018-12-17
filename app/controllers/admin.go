package controllers

import (
	"fmt"
	"net/http"
	"spi-web/app/models"

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
	err := models.CreateAdminUser()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
