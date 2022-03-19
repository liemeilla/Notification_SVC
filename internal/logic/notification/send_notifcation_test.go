package notification

import (
	"context"
	"reflect"
	"testing"

	"github.com/liemeilla/Notification_Service/internal/constant"
	"github.com/liemeilla/Notification_Service/internal/entities"
	testutil "github.com/liemeilla/Notification_Service/test-util"
)

func init() {
	testutil.InitDBAndExternalStub()
}

func TestSendNotification(t *testing.T) {
	ctx := testutil.InitContextForTest()

	type args struct {
		c             context.Context
		req           ReqSendNotification
		idempotencyID string
	}
	tests := []struct {
		name    string
		args    args
		wantRes ResSendNotification
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				c: ctx,
				req: ReqSendNotification{
					CustomerID: 21314,
					Mode:       "test",
					Notification: entities.Notification{
						ReferenceID:   "reference_id_1",
						TransactionID: "trn_id",
						PaymentStatus: "SUCCESS",
						Currency:      "IDR",
						Amount:        150000,
						ChannelCode:   "SHOPEEPAY",
					},
				},
				idempotencyID: "idempotency_id_1",
			},
			wantRes: ResSendNotification{
				Status: constant.STATUS_SUCCESS,
			},
		},
		{
			name: "failed case - customer url unreachable",
			args: args{
				c: ctx,
				req: ReqSendNotification{
					CustomerID: 21314,
					Mode:       "test",
					Notification: entities.Notification{
						ReferenceID:   "reference_id_2",
						TransactionID: "trn_id",
						PaymentStatus: "SUCCESS",
						Currency:      "IDR",
						Amount:        150000,
						ChannelCode:   "SHOPEEPAY",
					},
				},
				idempotencyID: "idempotency_id_1",
			},
			wantRes: ResSendNotification{
				Status: constant.STATUS_FAILED,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := SendNotification(tt.args.c, tt.args.req, tt.args.idempotencyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendNotification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("SendNotification() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
