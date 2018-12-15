package main

import (
	"fmt"
	"log"
	"spi-web/app/controllers"
	"spi-web/app/models"
	"spi-web/config"
	"spi-web/utils"

	"github.com/labstack/echo"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	defer models.Db.Close()
	e := echo.New()

	adminGroup := e.Group("/admin")
	adminGroup.Use(controllers.AdminMiddleWare)

	if err := e.Start(fmt.Sprintf(":%d", config.Config.Port)); err != nil {
		log.Fatalf("ListenAndServe err=%s", err)
	}
}
