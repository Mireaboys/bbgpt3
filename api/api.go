package api

import (
	"errors"
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Mess struct {
	Message string
	Error   error
}

type Answer struct {
	Text   string `json:"text" xml:"text" form:"text" query:"text"`
	Secret string `json:"secret" xml:"secret" form:"secret" query:"secret"`
	UUID   string `json:"uuid" xml:"uuid" form:"uuid" query:"uuid"`
}

func RunApi(port, tokenApi, tokenBot, secret string) {
	server := echo.New()
	gpt := NewGPT(tokenApi)
	bot := NewBot(tokenBot)
	server.Debug = true
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Mess{
			Message: "Service is available",
			Error:   nil,
		})
	})
	server.POST("/", func(c echo.Context) error {
		var ans Answer
		if err := c.Bind(&ans); err != nil {
			return c.JSON(http.StatusBadRequest, Mess{
				Message: "expected: {'text': 'string', 'secret': 'string'}",
				Error:   err,
			})
		}
		if secret != ans.Secret {
			return c.JSON(http.StatusForbidden, Mess{
				Message: "wrong secret",
				Error:   errors.New("permission denied"),
			})
		}

		res, err := gpt.Get(ans.Text)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Mess{
				Message: res,
				Error:   err,
			})
		}
		if ans.UUID != "" {
			if err := bot.SendMessage(res, ans.UUID); err != nil {
				return c.JSON(http.StatusInternalServerError, Mess{
					Message: res,
					Error:   err,
				})
			}
		}
		return c.JSON(http.StatusOK, Mess{
			Message: res,
			Error:   nil,
		})
	})
	// return http.Server{
	// 	Addr:    fmt.Sprintf(":%s", port),
	// 	Handler: server,
	// }
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", port)))
}
