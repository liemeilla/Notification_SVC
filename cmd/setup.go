package cmd

import (
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/liemeilla/Notification_Service/config"
	logicauth "github.com/liemeilla/Notification_Service/internal/logic/authentication"
	repositories "github.com/liemeilla/Notification_Service/internal/repositories"
)

func InitDB() {
	// psql info
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)

	// open database
	dbConn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("cannot connect to PostgresSQL caused by: ", err)
		panic(err)
	}
	// ping database
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("cannot ping to PostgresSQL caused by: ", err)
		panic(err)
	}
	// init db layer
	repositories.InitDBLayer(dbConn)
	repositories.InitExternalLayer()
	log.Println("Successfuly connected to database...")
}

func AuthApiKey() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(apiKey, password string, c echo.Context) (bool, error) {

		valid, err := logicauth.CheckAPIKey(c.Request().Context(), apiKey)
		if err != nil {
			return valid, nil
		}

		return valid, nil
	})
}
