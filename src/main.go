package main

import (
	"fmt"
	"github.com/labstack/echo"
	awsInit "main/aws"
	"main/nShortURL"
	"net/http"
)

func main() {
	if err := awsInit.InitAws("ap-northeast-2"); err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := nShortURL.InitNShortURL(); err != nil {
		fmt.Println(err.Error())
		return
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
