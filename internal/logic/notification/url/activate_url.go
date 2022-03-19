package notificationurl

import (
	"context"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/repositories"
)

type (
	ReqActivateURL struct {
		CustomerID int `json:"customer_id"`
	}
)

func ActivateURL(c context.Context, req ReqActivateURL) (res ResRegisterURL, err error) {
	notification, err := repositories.DBLayer.CustNotiRepo.FindByCustID(c, req.CustomerID)
	if err != nil {
		log.Error("catch error when repositories.DBLayer.CustNotiRepo.FindByCustID caused by: ", err.Error())
		return
	}

	if len(notification) == 0 {
		err = errors.New("customer id not exists")
		return
	}

	err = repositories.DBLayer.CustNotiRepo.UpdateNotificationURL(
		c,
		notification[0].CustomerId,
		notification[0].NotificationURL,
		constant.STATUS_ACTIVE,
	)
	if err != nil {
		log.Error("catch error when repositories.DBLayer.CustNotiRepo.UpdateNotificationURL caused by: ", err.Error())
		return
	}

	res.NotificationUrl = notification[0].NotificationURL

	return
}
