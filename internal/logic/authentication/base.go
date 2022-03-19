package authentication

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/labstack/gommon/log"

	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/logic/utils"
	"github.com/liemeilla/Notification_Service/internal/repositories"
)

type (
	ReqGenerateAPIKey struct {
		CustomerID  int          `json:"customer_id"`
		Permissions []Permission `json:"permissions"`
	}

	Permission struct {
		Name   string `json:"name"`
		Action string `json:"action"`
	}

	ResGenerateAPIKey struct {
		APIKey string `json:"api_key"`
	}
)

func CheckAPIKey(c context.Context, apiKey string) (valid bool, err error) {
	decodedBytes, err := utils.DecodeBase64(apiKey)
	if err != nil {
		log.Error("catch error when decode base 64: ", err.Error())
		return
	}

	var info ReqGenerateAPIKey
	err = json.Unmarshal(decodedBytes, &info)
	if err != nil {
		log.Error("catch error when unmarshal: ", err.Error())
		return
	}

	res, err := repositories.DBLayer.AuthRepo.FindByCustID(c, info.CustomerID)
	if err != nil {
		log.Error("catch error when repositories.DBLayer.AuthRepo.FindByCustID: ", err.Error())
		return
	}

	// log.Infof("ReqGenerateAPIKey: %+v", info)
	// log.Infof("RES FindByCustID: %+v", res)
	// log.Infof("API Key: %+v", apiKey)

	for _, v := range info.Permissions {
		if v.Action == constant.PERMISSION_NAME && v.Name == constant.PERMISSION_ACTION {
			valid = true
			return
		}
	}

	if len(res) == 0 {
		err = errors.New("api key not exists")
		log.Error("catch error when check response db: ", err.Error())
		return
	}

	if res[0].ApiKey == apiKey {
		valid = true
	}

	return
}
