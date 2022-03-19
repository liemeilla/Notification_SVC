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

func TestUpdateURL(t *testing.T) {
	ctx := testutil.InitContextForTest()

	type args struct {
		c   context.Context
		req ReqRegisterURL
	}
	tests := []struct {
		name    string
		args    args
		wantRes ResRegisterURL
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				c: ctx,
				req: ReqRegisterURL{
					CustomerID:      21314,
					NotificationUrl: "http://abcd.com/receive-noti",
				},
			},
			wantRes: ResRegisterURL{
				NotificationUrl: "http://abcd.com/receive-noti",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := UpdateURL(tt.args.c, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("UpdateURL() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
