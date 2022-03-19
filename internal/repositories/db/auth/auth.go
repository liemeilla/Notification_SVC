package auth

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"
	"github.com/liemeilla/Notification_Service/internal/logic/utils"
)

type Auth interface {
	Create(c context.Context, row entities.Auth) error
	FindByCustID(c context.Context, customer_id int) ([]entities.Auth, error)
	UpdateStatus(c context.Context, status string, customer_id string) error
}

type AuthRepo struct {
	DB *sql.DB
}

var AutRepo = new(AuthRepo)

func NewDB(db *sql.DB) Auth {
	AutRepo.DB = db
	return AutRepo
}

func (d *AuthRepo) Create(c context.Context, row entities.Auth) (err error) {
	query := fmt.Sprintf("INSERT INTO %s VALUES ($1,$2,$3,$4)", constant.TABLE_NAME_AUTH)

	utils.LogQueryDB(query)

	params := []interface{}{
		row.CustomerId,
		row.ApiKey,
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
func (d *AuthRepo) FindByCustID(c context.Context, customer_id int) (res []entities.Auth, err error) {
	query := fmt.Sprintf("SELECT customer_id, api_key, status, created_at FROM %s WHERE customer_id = $1",
		constant.TABLE_NAME_AUTH)

	utils.LogQueryDB(query)

	rows, err := d.DB.Query(query, customer_id)
	if err != nil {
		log.Error("catch error query FindByCustID caused by: ", err.Error())
		return
	}

	defer rows.Close()
	for rows.Next() {
		row := entities.Auth{}
		err = rows.Scan(
			&row.CustomerId,
			&row.ApiKey,
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
func (d *AuthRepo) UpdateStatus(c context.Context, status string, customer_id string) (err error) {
	query := strings.Join([]string{
		fmt.Sprintf("UPDATE %s SET", constant.TABLE_NAME_AUTH),
		"status = ?",
		"WHERE customer_id = ?",
	}, " ")

	utils.LogQueryDB(query)

	params := []interface{}{
		status,
		customer_id,
	}
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(
		params...,
	)
	return
}
