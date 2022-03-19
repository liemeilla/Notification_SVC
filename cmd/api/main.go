package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/liemeilla/Notification_Service/cmd"
	"github.com/liemeilla/Notification_Service/config"
	ctrlAuth "github.com/liemeilla/Notification_Service/internal/controller/authentication"
	ctrlMock "github.com/liemeilla/Notification_Service/internal/controller/mock"
	ctrlNoti "github.com/liemeilla/Notification_Service/internal/controller/notification"
)

func main() {
	// init db & external layer
	cmd.InitDB()

	// http server
	e := echo.New()

	ctrlAuth.InitAuthenticationController(e)
	ctrlNoti.InitNotificationController(e, cmd.AuthApiKey())
	ctrlMock.InitMockController(e)

	// general middleware
	e.Use(middleware.Logger())

	// start http server
	e.Logger.Fatal(e.Start(config.HTTP_SERVER_PORT))
}
