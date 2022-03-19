package external

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/internal/entities"
	"github.com/liemeilla/Notification_Service/internal/logic/utils"
)

type HTTP interface {
	SendNotificationToCustomer(
		customer_id int,
		apiKey string,
		url string,
		reqBody entities.ReqPaymentNotification,
	) (res entities.ResPaymentNotification, err error)
}

type HTTPRepo struct{}

var HTTPRep = new(HTTPRepo)

func NewHTTPRepo() HTTP {
	return HTTPRep
}

func (h *HTTPRepo) SendNotificationToCustomer(
	customer_id int,
	apiKey string,
	url string,
	reqBody entities.ReqPaymentNotification,
) (
	res entities.ResPaymentNotification,
	err error,
) {
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	c := http.Client{Timeout: time.Duration(10) * time.Second}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Error("catch error when create http new request: ", err.Error())
		return
	}

	bytesApiKey, err := json.Marshal(apiKey)
	if err != nil {
		log.Error("catch error when marshal: ", err.Error())
		return
	}

	encodedAuthKey := utils.Encodebase64ToString(bytesApiKey)

	req.Header.Add("Authorization", strings.Join([]string{"Basic", encodedAuthKey}, " "))
	req.Header.Add("content-type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		log.Error("catch error when c.Do: ", err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("catch error read body request: ", err.Error())
		return
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Error("catch error unmarshal body request: ", err.Error())
		return
	}

	log.Info("RESPONSE FROM CUSTOMER URL: ", res)
	return
}
