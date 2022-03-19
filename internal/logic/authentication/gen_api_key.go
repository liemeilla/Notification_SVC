package authentication

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"
	"github.com/liemeilla/Notification_Service/internal/logic/utils"
	"github.com/liemeilla/Notification_Service/internal/repositories"
)

func GenerateAPIKey(c context.Context, req ReqGenerateAPIKey) (res ResGenerateAPIKey, err error) {
	// check customer id exists or not
	dbResult, err := repositories.DBLayer.AuthRepo.FindByCustID(c, req.CustomerID)
	if err != nil {
		log.Error("catch error when repositories.DBLayer.AuthRepo.FindByCustID caused by: ", err.Error())
		return
	}

	if len(dbResult) == 1 {
		err = errors.New("api key already generated")
		return
	}

	// validate permission
	for _, v := range req.Permissions {
		if v.Action != constant.PERMISSION_ACTION || v.Name != constant.PERMISSION_NAME {
			err = errors.New("permission not valid")
			return
		}
	}

	// generate api key
	reqBytes, err := json.Marshal(req)
	if err != nil {
		log.Error("catch error when marshal permissions caused by: ", err.Error())
		return
	}

	// encode base64
	encodedAPIKey := utils.Encodebase64ToString(reqBytes)

	// insert into customer notification
	err = repositories.DBLayer.AuthRepo.Create(c, entities.Auth{
		CustomerId: req.CustomerID,
		ApiKey:     encodedAPIKey,
		Status:     constant.STATUS_ACTIVE,
	})
	if err != nil {
		log.Error("catch error when repositories.DBLayer.AuthRepo.Create caused by: ", err.Error())
		return
	}

	res = ResGenerateAPIKey{
		APIKey: encodedAPIKey,
	}
	return
}
