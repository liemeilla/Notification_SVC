package notification

import (
	"context"
	"reflect"
	"testing"

	"github.com/liemeilla/Notification_Service/internal/constant"
	testutil "github.com/liemeilla/Notification_Service/test-util"
)

func init() {
	testutil.InitDBAndExternalStub()
}

func TestRetrySendNotification(t *testing.T) {
	ctx := testutil.InitContextForTest()

	type args struct {
		c             context.Context
		req           ReqRetrySendNotification
		idempotencyID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes ResSendNotification
		wantErr bool
	}{
		{
			name: "success case - retry to send notification",
			args: args{
				c: ctx,
				req: ReqRetrySendNotification{
					CustomerID: 21314,
				},
				idempotencyID: "idempotency_id_2",
			},
			wantRes: ResSendNotification{
				Status: constant.STATUS_SUCCESS,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := RetrySendNotification(tt.args.c, tt.args.req, tt.args.idempotencyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RetrySendNotification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("RetrySendNotification() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
