package notification

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"
	"github.com/liemeilla/Notification_Service/internal/repositories"
)

type (
	ReqSendNotification struct {
		CustomerID   int                   `json:"customer_id"`
		Mode         string                `json:"mode"`
		Notification entities.Notification `json:"notification"`
	}

	ResSendNotification struct {
		Status string `json:"status"`
	}
)

func SendNotification(c context.Context, req ReqSendNotification, idempotencyID string) (res ResSendNotification, err error) {

	fmt.Printf("ReqSendNotification: %+v ", req)
	fmt.Println()

	// check idempotency id, if exists return existing response
	notiLog, err := repositories.DBLayer.NotiLogRepo.FindByIdempotencyID(c, idempotencyID)
	if err != nil {
		log.Error("catch error when repositories.DBLayer.NotiLogRepo.FindByIdempotencyID caused by: ", err.Error())
		return
	}
	if len(notiLog) > 0 {
		// if exists in db, return the response
		err = json.Unmarshal([]byte(notiLog[0].ResponseJson), &res)
		if err != nil {
			log.Error("catch error when unmarshal response json caused by: ", err.Error())
			return
		}
		return
	}

	// get noti url & status
	notiInfo, err := repositories.DBLayer.CustNotiRepo.FindByCustID(c, req.CustomerID)
	if err != nil {
		log.Error("catch error when repositories.DBLayer.CustNotiRepo.FindByCustID caused by: ", err.Error())
		return
	}

	if len(notiInfo) == 0 {
		err = errors.New("customer id not exists")
		return
	}

	notiUrl := notiInfo[0].NotificationURL

	switch req.Mode {
	case constant.TEST_MODE:
		res, err = sendNotification(c, req, notiUrl, idempotencyID)
		if err != nil {
			return
		}
	case constant.REAL_MODE:
		// check status of url
		result, err := repositories.DBLayer.CustNotiRepo.FindByCustID(c, req.CustomerID)
		if err != nil {
			return res, err
		}

		if len(result) == 0 {
			err = errors.New("customer not exists")
			return res, err
		}

		if result[0].Status != constant.STATUS_ACTIVE {
			err = errors.New("registered url is not activated, please activate your registered url")
			return res, err
		}

		res, err = sendNotification(c, req, notiUrl, idempotencyID)
		if err != nil {
			return res, err
		}
	default:
		err = errors.New("mode invalid")
		return
	}

	return
}

func sendNotification(
	ctx context.Context,
	req ReqSendNotification,
	notificationUrl string,
	id string,
) (resp ResSendNotification, err error) {
	failedCounter := 0
	success := false

	var resNoti entities.ResPaymentNotification

	authInfo, err := repositories.DBLayer.AuthRepo.FindByCustID(ctx, req.CustomerID)
	if err != nil {
		return
	}

	if len(authInfo) == 0 {
		err = errors.New("customer id not exsits")
		return
	}

	for failedCounter < 3 && !success {
		resNoti, err = repositories.ExternalLayer.HTTPRepo.SendNotificationToCustomer(
			req.CustomerID,
			authInfo[0].ApiKey,
			notificationUrl,
			entities.ReqPaymentNotification{
				Notification: req.Notification,
			},
		)
		if err != nil || resNoti.Status != constant.STATUS_SUCCESS {
			failedCounter++
			log.Error("failed to send notification count-", failedCounter)

			// sleep for 5 seconds
			time.Sleep(time.Second * 5)
			continue
		}

		success = true
	}

	statusSent := constant.STATUS_FAILED
	if err == nil && success {
		statusSent = constant.STATUS_SUCCESS
	}

	bytesData, err := json.Marshal(req.Notification)
	if err != nil {
		log.Error("catch error when marshal req.Notification: ", err.Error())
		return
	}

	resp.Status = statusSent

	reqJson, err := json.Marshal(req)
	if err != nil {
		log.Error("catch error when marshal req json: ", err.Error())
		return
	}
	resJson, err := json.Marshal(resp)
	if err != nil {
		log.Error("catch error when marshal res json: ", err.Error())
		return
	}

	// insert into noti log
	err = repositories.DBLayer.NotiLogRepo.Create(context.Background(), entities.NotificationLog{
		IdempotencyId:    id,
		CustomerId:       req.CustomerID,
		NotificationUrl:  notificationUrl,
		NotificationData: string(bytesData),
		RequestJson:      string(reqJson),
		ResponseJson:     string(resJson),
		StatusSent:       statusSent,
	})
	if err != nil {
		log.Error("catch error when repositories.DBLayer.NotiLogRepo.Create caused by: ", err.Error())
		return
	}

	return
}
