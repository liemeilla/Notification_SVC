package repositories

import (
	"database/sql"

	"github.com/liemeilla/Notification_Service/internal/repositories/db/auth"
	customernotification "github.com/liemeilla/Notification_Service/internal/repositories/db/customer-notification"
	notificationlog "github.com/liemeilla/Notification_Service/internal/repositories/db/notification-log"

	authstub "github.com/liemeilla/Notification_Service/internal/repositories/db/auth/stub"
	custnotistub "github.com/liemeilla/Notification_Service/internal/repositories/db/customer-notification/stub"
	notilogstub "github.com/liemeilla/Notification_Service/internal/repositories/db/notification-log/stub"
)

type DB struct {
	AuthRepo     auth.Auth
	CustNotiRepo customernotification.CustomerNotification
	NotiLogRepo  notificationlog.NotificationLog
}

var DBLayer DB

func InitDBLayer(db *sql.DB) {
	DBLayer.AuthRepo = auth.NewDB(db)
	DBLayer.CustNotiRepo = customernotification.NewDB(db)
	DBLayer.NotiLogRepo = notificationlog.NewDB(db)
}

func InitDBLayerStub() {
	DBLayer.AuthRepo = authstub.NewStubDB()
	DBLayer.CustNotiRepo = custnotistub.NewStubDB()
	DBLayer.NotiLogRepo = notilogstub.NewStubDB()
}
