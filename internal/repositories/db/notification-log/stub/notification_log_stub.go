package stub

import (
	"context"
	"encoding/json"
	"time"

	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"
	notificationlog "github.com/liemeilla/Notification_Service/internal/repositories/db/notification-log"
)

type NotificationLogRepoStub struct {
}

var NotiLogRepoStub = new(NotificationLogRepoStub)

func NewStubDB() notificationlog.NotificationLog {
	return NotiLogRepoStub
}

const IDEMPOTENCY_ID_2 = "idempotency_id_2"

func (d *NotificationLogRepoStub) Create(c context.Context, row entities.NotificationLog) (err error) {

	return
}
func (d *NotificationLogRepoStub) FindByIdempotencyID(c context.Context, id string) (res []entities.NotificationLog, err error) {

	if id == IDEMPOTENCY_ID_2 {
		res = []entities.NotificationLog{
			{
				IdempotencyId:    "idempotency_id_2",
				CustomerId:       21314,
				NotificationUrl:  "http://abc.com/receive-noti",
				NotificationData: getNotificationData("reference_id_1"),
				RequestJson:      "{\"customer_id\":5667,\"mode\":\"test\",\"notification\":{\"reference_id\":\"reference_id_1\",\"transaction_id\":\"trn_id_1\",\"transaction_time\":\"2017-07-21T17:32:28Z\",\"payment_status\":\"SUCCESS\",\"currency\":\"IDR\",\"amount\":1000,\"channel_code\":\"SHOPEEPAY\"}}",
				ResponseJson:     "{\"status\":\"failed\"}",
				StatusSent:       constant.STATUS_FAILED,
			},
		}
	}

	return
}

func (d *NotificationLogRepoStub) UpdateStatusAndResponse(c context.Context, status string, id string, respJson string) (err error) {

	return
}

func getNotificationData(refId string) string {
	var notiData = []entities.Notification{
		{
			ReferenceID:     "reference_id_1",
			TransactionID:   "trn_id_1",
			TransactionTime: time.Now(),
			PaymentStatus:   "SUCCESS",
			Currency:        "IDR",
			Amount:          1000,
			ChannelCode:     "SHOPEEPAY",
		},
	}
	var foundedData entities.Notification
	for _, v := range notiData {
		if v.ReferenceID == refId {
			foundedData = v
			break
		}
	}
	bytes, _ := json.Marshal(foundedData)
	return string(bytes)
}
