package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	awsInit "main/aws"
	"main/nShortURL"
	"main/validator"
	"net/http"
)

type ReqGetShortUrl struct {
	Url string `query:"url"`
}

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
	e.GET("/shortUrl", func(c echo.Context) error {
		req := &ReqGetShortUrl{}
		if err := validator.ValidateReq(c, req); err != nil {
			return err
		}
		result, err := nShortURL.GetShortURL(req.Url)
		fmt.Println("여기 들어오지?")
		fmt.Println(result)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, result.Result.Url)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
