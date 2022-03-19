package stub

import (
	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"
	external "github.com/liemeilla/Notification_Service/internal/repositories/external"
)

type HTTPRepoStub struct{}

var HTTPRepStub = new(HTTPRepoStub)

func NewHTTPRepoStub() external.HTTP {
	return HTTPRepStub
}

func (h *HTTPRepoStub) SendNotificationToCustomer(
	customer_id int,
	apiKey string,
	url string,
	reqBody entities.ReqPaymentNotification,
) (
	res entities.ResPaymentNotification,
	err error,
) {
	if reqBody.Notification.ReferenceID == "reference_id_1" {
		res.Status = constant.STATUS_SUCCESS
	} else if reqBody.Notification.ReferenceID == "reference_id_2" {
		res.Status = constant.STATUS_FAILED
	}

	return
}
