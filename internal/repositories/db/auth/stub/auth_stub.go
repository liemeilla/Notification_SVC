package stub

import (
	"context"

	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"
	"github.com/liemeilla/Notification_Service/internal/repositories/db/auth"
)

type AuthRepoStub struct {
}

var AutRepoStub = new(AuthRepoStub)

func NewStubDB() auth.Auth {
	return AutRepoStub
}

const (
	CUST_ID_21314 = 21314
)

func (d *AuthRepoStub) Create(c context.Context, row entities.Auth) (err error) {

	return
}
func (d *AuthRepoStub) FindByCustID(c context.Context, customer_id int) (res []entities.Auth, err error) {

	if customer_id == CUST_ID_21314 {
		res = []entities.Auth{
			{
				CustomerId: customer_id,
				ApiKey:     "eyJjdXN0b21lcl9pZCI6MjEzMTQsInBlcm1pc3Npb25zIjpbeyJuYW1lIjoibm90aWZpY2F0aW9uIiwiYWN0aW9uIjoid3JpdGUifV19",
				Status:     constant.STATUS_ACTIVE,
			},
		}
	}

	return
}
func (d *AuthRepoStub) UpdateStatus(c context.Context, status string, customer_id string) (err error) {
	return
}
