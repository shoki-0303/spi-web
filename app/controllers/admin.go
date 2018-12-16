package controllers

import (
	"fmt"
	"spi-web/app/models"

	"github.com/labstack/echo"
)

// AdminMiddleWare : done before processeing related to admin
func AdminMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	//"ここにadminUserを確かめる処理を書く"
	return next
}

func AdminCreateUser(c echo.Context) error {
	err := models.CreateAdminUser()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
