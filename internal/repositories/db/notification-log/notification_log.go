package notificationlog

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"
	"github.com/liemeilla/Notification_Service/internal/logic/utils"
)

type NotificationLog interface {
	Create(c context.Context, row entities.NotificationLog) error
	FindByIdempotencyID(c context.Context, id string) ([]entities.NotificationLog, error)
	UpdateStatusAndResponse(c context.Context, status string, id string, respJson string) error
}

type NotificationLogRepo struct {
	DB *sql.DB
}

var NotiLogRepo = new(NotificationLogRepo)

func NewDB(db *sql.DB) NotificationLog {
	NotiLogRepo.DB = db
	return NotiLogRepo
}

func (d *NotificationLogRepo) Create(c context.Context, row entities.NotificationLog) (err error) {
	query := fmt.Sprintf("INSERT INTO %s VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", constant.TABLE_NAME_NOTIFICATION_LOG)

	utils.LogQueryDB(query)

	params := []interface{}{
		row.IdempotencyId,
		row.CustomerId,
		row.NotificationUrl,
		row.NotificationData,
		row.RequestJson,
		row.ResponseJson,
		row.StatusSent,
		"NOW()",
	}
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return
	}
	_, err = stmt.Exec(params...)

	return
}
func (d *NotificationLogRepo) FindByIdempotencyID(c context.Context, id string) (res []entities.NotificationLog, err error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1",
		constant.TABLE_NAME_NOTIFICATION_LOG)

	utils.LogQueryDB(query)

	rows, err := d.DB.Query(query, id)
	if err != nil {
		log.Error("catch error query FindByCustID caused by: ", err.Error())
		return
	}

	defer rows.Close()
	for rows.Next() {
		row := entities.NotificationLog{}
		err = rows.Scan(
			&row.IdempotencyId,
			&row.CustomerId,
			&row.NotificationUrl,
			&row.NotificationData,
			&row.RequestJson,
			&row.ResponseJson,
			&row.StatusSent,
			&row.CreatedAt,
		)
		if err != nil {
			return
		}
		res = append(res, row)
	}

	return
}

func (d *NotificationLogRepo) UpdateStatusAndResponse(c context.Context, status string, id string, respJson string) (err error) {
	query := fmt.Sprintf("UPDATE %s SET status_sent = $1, response_json = $2 WHERE id = $3", constant.TABLE_NAME_NOTIFICATION_LOG)

	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return
	}
	_, err = stmt.Exec(status, respJson, id)

	return
}
