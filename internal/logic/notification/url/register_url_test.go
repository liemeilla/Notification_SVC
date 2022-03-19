package notificationurl

import (
	"context"
	"reflect"
	"testing"

	testutil "github.com/liemeilla/Notification_Service/test-util"
)

func init() {
	testutil.InitDBAndExternalStub()
}

func TestRegisterURL(t *testing.T) {
	ctx := testutil.InitContextForTest()

	type args struct {
		c   context.Context
		req ReqRegisterURL
	}
	tests := []struct {
		name     string
		args     args
		wantResp ResRegisterURL
		wantErr  bool
	}{
		{
			name: "success case",
			args: args{
				c: ctx,
				req: ReqRegisterURL{
					CustomerID:      1234,
					NotificationUrl: "http://xendit.com/receive-noti",
				},
			},
			wantResp: ResRegisterURL{
				NotificationUrl: "http://xendit.com/receive-noti",
			},
			wantErr: false,
		},
		{
			name: "failed case",
			args: args{
				c: ctx,
				req: ReqRegisterURL{
					CustomerID:      21314,
					NotificationUrl: "http://xendit.com/receive-noti",
				},
			},
			wantResp: ResRegisterURL{
				NotificationUrl: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := RegisterURL(tt.args.c, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RegisterURL() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
