package notificationurl

import (
	"context"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"
	"github.com/liemeilla/Notification_Service/internal/repositories"
)

type (
	ReqRegisterURL struct {
		CustomerID      int    `json:"customer_id"`
		NotificationUrl string `json:"notification_url"`
	}

	ResRegisterURL struct {
		NotificationUrl string `json:"notification_url"`
	}
)

func RegisterURL(c context.Context, req ReqRegisterURL) (resp ResRegisterURL, err error) {
	res, err := repositories.DBLayer.CustNotiRepo.FindByCustID(c, req.CustomerID)
	if err != nil {
		log.Error("catch error when repositories.DBLayer.CustNotiRepo.FindByCustID caused by: ", err.Error())
		return
	}

	if len(res) == 1 {
		err = errors.New("notification url already registered")
		return
	}

	err = repositories.DBLayer.CustNotiRepo.Create(c, entities.CustomerNotification{
		CustomerId:      req.CustomerID,
		NotificationURL: req.NotificationUrl,
		Status:          constant.STATUS_PENDING,
	})
	if err != nil {
		log.Error("catch error when repositories.DBLayer.CustNotiRepo.Create caused by: ", err.Error())
		return
	}

	resp.NotificationUrl = req.NotificationUrl

	return
}
