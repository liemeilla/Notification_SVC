package notification

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"
	"github.com/liemeilla/Notification_Service/internal/repositories"
)

type (
	ReqRetrySendNotification struct {
		CustomerID int `json:"customer_id"`
	}
)

func RetrySendNotification(c context.Context, req ReqRetrySendNotification, idempotencyID string) (res ResSendNotification, err error) {
	// get notification data from noti log by idempotency id
	notiLog, err := repositories.DBLayer.NotiLogRepo.FindByIdempotencyID(c, idempotencyID)
	if err != nil {
		log.Error("catch error when repositories.DBLayer.NotiLogRepo.FindByIdempotencyID caused by: ", err.Error())
		return
	}

	if len(notiLog) > 0 {
		if notiLog[0].StatusSent == constant.STATUS_SUCCESS {
			err = json.Unmarshal([]byte(notiLog[0].ResponseJson), &res)
			if err != nil {
				log.Error("catch error when unmarshal response json caused by: ", err.Error())
				return
			}
			return
		} else { // retry send noti
			var notificationData entities.ReqPaymentNotification
			err = json.Unmarshal([]byte(notiLog[0].NotificationData), &notificationData.Notification)
			if err != nil {
				return
			}
			// get noti url & status
			notiInfo, err := repositories.DBLayer.CustNotiRepo.FindByCustID(c, req.CustomerID)
			if err != nil {
				log.Error("catch error when repositories.DBLayer.CustNotiRepo.FindByCustID caused by: ", err.Error())
				return res, err
			}

			if len(notiInfo) == 0 {
				err = errors.New("customer id not exists")
				return res, err
			}

			// get api key
			authInfo, err := repositories.DBLayer.AuthRepo.FindByCustID(c, req.CustomerID)
			if err != nil {
				return res, err
			}

			if len(authInfo) == 0 {
				err = errors.New("customer id not exsits")
				return res, err
			}

			// send to customer url
			notiUrl := notiInfo[0].NotificationURL
			resNoti, err := repositories.ExternalLayer.HTTPRepo.SendNotificationToCustomer(
				req.CustomerID,
				authInfo[0].ApiKey,
				notiUrl,
				notificationData,
			)
			if err != nil {
				return res, err
			}

			// update status to be success and response json
			if resNoti.Status == constant.STATUS_SUCCESS {
				respJson, err := json.Marshal(resNoti)
				if err != nil {
					return res, err
				}
				err = repositories.DBLayer.NotiLogRepo.UpdateStatusAndResponse(
					c,
					constant.STATUS_SUCCESS,
					idempotencyID,
					string(respJson),
				)
				if err != nil {
					return res, err
				}
			}

			res.Status = resNoti.Status
		}
	} else {
		err = errors.New("notification not exists")
		return
	}

	return
}
