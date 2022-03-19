package stub

import (
	"context"

	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"

	customernotification "github.com/liemeilla/Notification_Service/internal/repositories/db/customer-notification"
)

type CustomerNotificationRepoStub struct{}

var CustNotiRepoStub = new(CustomerNotificationRepoStub)

func NewStubDB() customernotification.CustomerNotification {
	return CustNotiRepoStub
}

const (
	CUST_ID_21314 = 21314
)

func (d *CustomerNotificationRepoStub) Create(c context.Context, row entities.CustomerNotification) (err error) {

	return
}

func (d *CustomerNotificationRepoStub) FindByCustID(c context.Context, customer_id int) (res []entities.CustomerNotification, err error) {
	if customer_id == CUST_ID_21314 {
		res = []entities.CustomerNotification{
			{
				CustomerId:      customer_id,
				NotificationURL: "http://abc.com/receive-noti",
				Status:          constant.STATUS_PENDING,
			},
		}
	}
	return
}
func (d *CustomerNotificationRepoStub) UpdateNotificationURL(
	c context.Context,
	customer_id int,
	url string,
	status string,
) (err error) {

	return
}
