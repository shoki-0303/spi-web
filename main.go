package main

import (
	"fmt"
	"log"
	"spi-web/config"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	if err := e.Start(fmt.Sprintf(":%d", config.Config.Port)); err != nil {
		log.Fatalf("ListenAndServe err=%s", err)
	}
}
