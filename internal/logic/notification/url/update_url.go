package notificationurl

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/repositories"
)

func UpdateURL(c context.Context, req ReqRegisterURL) (res ResRegisterURL, err error) {
	err = repositories.DBLayer.CustNotiRepo.UpdateNotificationURL(
		c,
		req.CustomerID,
		req.NotificationUrl,
		constant.STATUS_PENDING,
	)
	if err != nil {
		log.Error("catch error when repositories.DBLayer.CustNotiRepo.Create caused by: ", err.Error())
		return
	}
	res.NotificationUrl = req.NotificationUrl
	return
}
