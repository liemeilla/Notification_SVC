package testutil

import (
	"context"

	"github.com/liemeilla/Notification_Service/internal/repositories"
)

func InitDBAndExternalStub() {
	repositories.InitDBLayerStub()
	repositories.InitExternalLayerStub()
}

func InitContextForTest() context.Context {
	newCtx := context.Background()
	return newCtx
}
