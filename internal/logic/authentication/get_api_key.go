package authentication

import (
	"context"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/internal/repositories"
)

func GetAPIKey(c context.Context, customer_id int) (res ResGenerateAPIKey, err error) {
	dbResult, err := repositories.DBLayer.AuthRepo.FindByCustID(c, customer_id)
	if err != nil {
		log.Error("catch error when repositories.DBLayer.AuthRepo.FindByCustID caused by: ", err.Error())
		return
	}

	if len(dbResult) == 0 {
		err = errors.New("api key not exists")
		return
	}

	res = ResGenerateAPIKey{
		APIKey: dbResult[0].ApiKey,
	}
	return
}
