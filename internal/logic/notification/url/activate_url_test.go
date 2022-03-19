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

func TestActivateURL(t *testing.T) {
	ctx := testutil.InitContextForTest()

	type args struct {
		c   context.Context
		req ReqActivateURL
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
				req: ReqActivateURL{
					CustomerID: 21314,
				},
			},
			wantRes: ResRegisterURL{
				NotificationUrl: "http://abc.com/receive-noti",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := ActivateURL(tt.args.c, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ActivateURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ActivateURL() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
