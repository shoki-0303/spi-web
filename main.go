package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"spi-web/app/controllers"
	"spi-web/app/controllers/helpers"
	"spi-web/app/models"
	"spi-web/config"
	"spi-web/utils"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var t = &Template{
	templates: template.Must(template.ParseGlob("app/views/*.html")),
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = http.StatusText(code)
	}
	data := map[string]string{
		"code": strconv.Itoa(code),
		"msg":  msg,
	}
	c.Render(code, "error.html", data)
	c.Logger().Error(err)
}

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	defer models.Db.Close()

	e := echo.New()
	e.Pre(helpers.MethodOverride)
	e.Static("admin/public/scss", "./app/public/scss")
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Renderer = t
	e.Use(middleware.Logger())

	adminGroup := e.Group("/admin")
	adminGroup.Use(controllers.AdminMiddleWare)
	adminGroup.GET("/register", controllers.AdminRegister)
	adminGroup.GET("/login", controllers.AdminLogin)
	adminGroup.POST("/user", controllers.AdminCreateUser)
	adminGroup.POST("/login", controllers.ConfirmAdminUser)

	restrected := adminGroup.Group("/restricted")
	restrected.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "query:token",
	}))
	restrected.GET("/:name", controllers.ShowAdminUser)
	restrected.GET("/update", controllers.AdminUpdate)
	restrected.PATCH("/update", controllers.UpdateAdminUser)

	if err := e.Start(fmt.Sprintf(":%d", config.Config.Port)); err != nil {
		log.Fatalf("ListenAndServe err=%s", err)
	}
}
