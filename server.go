package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/mrtc0/csp-server/report"
	"github.com/mrtc0/csp-server/slack"
)

func main() {
	e := echo.New()
	e.POST("/report", func(c echo.Context) error {
		report := new(report.Report)
		if err := c.Bind(report); err != nil {
			return c.JSON(http.StatusBadRequest, report)
		}
		slack.Send(*report)
		return c.JSON(http.StatusOK, report)
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
