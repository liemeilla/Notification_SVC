package customernotification

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"
	"github.com/liemeilla/Notification_Service/internal/logic/utils"
)

type CustomerNotification interface {
	Create(c context.Context, row entities.CustomerNotification) error
	FindByCustID(c context.Context, customer_id int) ([]entities.CustomerNotification, error)
	UpdateNotificationURL(c context.Context, customer_id int, url string, status string) error
}

type CustomerNotificationRepo struct {
	DB *sql.DB
}

var CustNotiRepo = new(CustomerNotificationRepo)

func NewDB(db *sql.DB) CustomerNotification {
	CustNotiRepo.DB = db
	return CustNotiRepo
}

func (d *CustomerNotificationRepo) Create(c context.Context, row entities.CustomerNotification) (err error) {
	query := fmt.Sprintf("INSERT INTO %s VALUES ($1,$2,$3,$4)", constant.TABLE_NAME_CUSTOMER_NOTIFICATION)

	utils.LogQueryDB(query)

	params := []interface{}{
		row.CustomerId,
		row.NotificationURL,
		row.Status,
		"NOW()",
	}
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return
	}
	_, err = stmt.Exec(params...)

	return
}

func (d *CustomerNotificationRepo) FindByCustID(c context.Context, customer_id int) (res []entities.CustomerNotification, err error) {
	query := fmt.Sprintf("SELECT customer_id, notification_url, status, created_at FROM %s WHERE customer_id = $1",
		constant.TABLE_NAME_CUSTOMER_NOTIFICATION)

	utils.LogQueryDB(query)

	rows, err := d.DB.Query(query, customer_id)
	if err != nil {
		log.Error("catch error query FindByCustID caused by: ", err.Error())
		return
	}

	defer rows.Close()
	for rows.Next() {
		row := entities.CustomerNotification{}
		err = rows.Scan(
			&row.CustomerId,
			&row.NotificationURL,
			&row.Status,
			&row.CreatedAt,
		)
		if err != nil {
			return
		}
		res = append(res, row)
	}

	return
}
func (d *CustomerNotificationRepo) UpdateNotificationURL(
	c context.Context,
	customer_id int,
	url string,
	status string,
) (err error) {
	query := fmt.Sprintf("UPDATE %s SET status = $1, notification_url =$2 WHERE customer_id = $3", constant.TABLE_NAME_CUSTOMER_NOTIFICATION)

	utils.LogQueryDB(query)

	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return
	}
	_, err = stmt.Exec(status, url, customer_id)

	return
}
