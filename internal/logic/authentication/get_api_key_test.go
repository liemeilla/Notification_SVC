package authentication

import (
	"context"
	"reflect"
	"testing"

	testutil "github.com/liemeilla/Notification_Service/test-util"
)

func init() {
	testutil.InitDBAndExternalStub()
}

func TestGetAPIKey(t *testing.T) {
	ctx := testutil.InitContextForTest()

	type args struct {
		c           context.Context
		customer_id int
	}
	tests := []struct {
		name    string
		args    args
		wantRes ResGenerateAPIKey
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				c:           ctx,
				customer_id: 21314,
			},
			wantRes: ResGenerateAPIKey{
				APIKey: "eyJjdXN0b21lcl9pZCI6MjEzMTQsInBlcm1pc3Npb25zIjpbeyJuYW1lIjoibm90aWZpY2F0aW9uIiwiYWN0aW9uIjoid3JpdGUifV19",
			},
			wantErr: false,
		},
		{
			name: "failed case",
			args: args{
				c:           ctx,
				customer_id: 123,
			},
			wantRes: ResGenerateAPIKey{
				APIKey: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := GetAPIKey(tt.args.c, tt.args.customer_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetAPIKey() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
