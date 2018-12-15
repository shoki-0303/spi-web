package main

import (
	"fmt"
	"log"
	"spi-web/config"
	"spi-web/utils"

	"github.com/labstack/echo"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	e := echo.New()

	if err := e.Start(fmt.Sprintf(":%d", config.Config.Port)); err != nil {
		log.Fatalf("ListenAndServe err=%s", err)
	}
}
