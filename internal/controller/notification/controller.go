package notification

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/config"
	"github.com/liemeilla/Notification_Service/internal/logic/notification"
	notificationurl "github.com/liemeilla/Notification_Service/internal/logic/notification/url"
	"github.com/liemeilla/Notification_Service/internal/logic/utils"
)

var PrefixNotification = "/notification"
var PrefixURL = "/url"

func InitNotificationController(e *echo.Echo, middlewares ...echo.MiddlewareFunc) {
	e.POST(config.PREFIX_API+PrefixNotification, SendNotification, middlewares...)
	e.POST(config.PREFIX_API+PrefixNotification+"/retry", RetrySendNotification, middlewares...)

	e.POST(config.PREFIX_API+PrefixNotification+PrefixURL+"/register", RegisterURL, middlewares...)
	e.POST(config.PREFIX_API+PrefixNotification+PrefixURL+"/activate", ActivateURL, middlewares...)
	e.POST(config.PREFIX_API+PrefixNotification+PrefixURL+"/update", UpdateURL, middlewares...)
}

func SendNotification(c echo.Context) (err error) {
	var (
		req notification.ReqSendNotification
		res notification.ResSendNotification
	)

	err = c.Bind(&req)
	if err != nil {
		log.Error("catch error when c.Bind caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	idempotencyID := c.Request().Header.Get("X-Idempotency-Key")
	res, err = notification.SendNotification(c.Request().Context(), req, idempotencyID)
	if err != nil {
		log.Error("catch error when notification.SendNotification caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}

func RetrySendNotification(c echo.Context) (err error) {
	var (
		req notification.ReqRetrySendNotification
		res notification.ResSendNotification
	)
	err = c.Bind(&req)
	if err != nil {
		log.Error("catch error when c.Bind caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	idempotencyID := c.Request().Header.Get("X-Idempotency-Key")
	res, err = notification.RetrySendNotification(c.Request().Context(), req, idempotencyID)
	if err != nil {
		log.Error("catch error when notification.RetrySendNotification caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}
	return c.JSON(http.StatusOK, res)
}

func RegisterURL(c echo.Context) (err error) {
	var (
		req notificationurl.ReqRegisterURL
		res notificationurl.ResRegisterURL
	)
	err = c.Bind(&req)
	if err != nil {
		log.Error("catch error when c.Bind caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}
	res, err = notificationurl.RegisterURL(c.Request().Context(), req)
	if err != nil {
		log.Error("catch error when notificationurl.ReqRegisterURL caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateURL(c echo.Context) (err error) {
	var (
		req notificationurl.ReqRegisterURL
		res notificationurl.ResRegisterURL
	)

	err = c.Bind(&req)
	if err != nil {
		log.Error("catch error when c.Bind caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	res, err = notificationurl.UpdateURL(c.Request().Context(), req)
	if err != nil {
		log.Error("catch error when notificationurl.UpdateURL caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}

func ActivateURL(c echo.Context) (err error) {
	var (
		req notificationurl.ReqActivateURL
		res notificationurl.ResRegisterURL
	)

	err = c.Bind(&req)
	if err != nil {
		log.Error("catch error when c.Bind caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	res, err = notificationurl.ActivateURL(c.Request().Context(), req)
	if err != nil {
		log.Error("catch error when notificationurl.ActivateURL caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}
