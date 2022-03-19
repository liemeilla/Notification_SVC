package repositories

import (
	"github.com/liemeilla/Notification_Service/internal/repositories/external"
	externalstub "github.com/liemeilla/Notification_Service/internal/repositories/external/stub"
)

type (
	External struct {
		HTTPRepo external.HTTP
	}
)

var ExternalLayer External

func InitExternalLayer() {
	ExternalLayer.HTTPRepo = external.NewHTTPRepo()
}

func InitExternalLayerStub() {
	ExternalLayer.HTTPRepo = externalstub.NewHTTPRepoStub()
}
