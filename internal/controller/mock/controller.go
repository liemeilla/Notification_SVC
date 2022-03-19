package mock

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/liemeilla/Notification_Service/internal/entities"
)

const (
	PrefixMock1 = "/mock1"
	PrefixMock2 = "/mock2"
	Endpoint    = "/receive-notification"
)

func InitMockController(e *echo.Echo) {
	e.POST(PrefixMock1+Endpoint, Mock_1) // success
	e.POST(PrefixMock2+Endpoint, Mock_2) // failed
}

func Mock_1(c echo.Context) (err error) {
	var (
		req entities.ReqPaymentNotification
		res entities.ResPaymentNotification
	)

	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	res.Status = "success"
	return c.JSON(http.StatusOK, res)
}

func Mock_2(c echo.Context) (err error) {
	var (
		req entities.ReqPaymentNotification
		res entities.ResPaymentNotification
	)

	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	res.Status = "failed"
	return c.JSON(http.StatusOK, res)
}
