package notification

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/config"
	logicauth "github.com/liemeilla/Notification_Service/internal/logic/authentication"
	"github.com/liemeilla/Notification_Service/internal/logic/utils"
)

type AuthenticationController struct{}

var PrefixAuth = "/authentication"

func InitAuthenticationController(e *echo.Echo, middlewares ...echo.MiddlewareFunc) {
	authCtrl := AuthenticationController{}

	e.POST(config.PREFIX_API+PrefixAuth+"/generate-api-key", authCtrl.GenerateAPIKey, middlewares...)
	e.GET(config.PREFIX_API+PrefixAuth+"/get-api-key/:cust_id", authCtrl.GetAPIKey, middlewares...)
}

func (ctrl *AuthenticationController) GenerateAPIKey(c echo.Context) (err error) {
	var (
		reqGenAPIKey logicauth.ReqGenerateAPIKey
		res          logicauth.ResGenerateAPIKey
	)

	err = c.Bind(&reqGenAPIKey)
	if err != nil {
		log.Error("catch error when c.Bind caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	res, err = logicauth.GenerateAPIKey(c.Request().Context(), reqGenAPIKey)
	if err != nil {
		log.Error("catch error when logicauth.GenerateAPIKey caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}

func (ctrl *AuthenticationController) GetAPIKey(c echo.Context) (err error) {
	var (
		res logicauth.ResGenerateAPIKey
	)

	customerID := c.Param("cust_id")
	customerIDNum, err := strconv.Atoi(customerID)
	if err != nil {
		log.Error("catch error when convert customer id caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	res, err = logicauth.GetAPIKey(c.Request().Context(), customerIDNum)
	if err != nil {
		log.Error("catch error when logicauth.GenerateAPIKey caused by: ", err.Error())
		return c.JSON(http.StatusBadRequest, utils.MappingErrorResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}
