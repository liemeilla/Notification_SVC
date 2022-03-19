package authentication

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

func TestGenerateAPIKey(t *testing.T) {
	ctx := testutil.InitContextForTest()

	type args struct {
		c   context.Context
		req ReqGenerateAPIKey
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
				c: ctx,
				req: ReqGenerateAPIKey{
					CustomerID: 5667,
					Permissions: []Permission{
						{
							Name:   constant.PERMISSION_NAME,
							Action: constant.PERMISSION_ACTION,
						},
					},
				},
			},
			wantRes: ResGenerateAPIKey{
				APIKey: "eyJjdXN0b21lcl9pZCI6NTY2NywicGVybWlzc2lvbnMiOlt7Im5hbWUiOiJub3RpZmljYXRpb24iLCJhY3Rpb24iOiJ3cml0ZSJ9XX0=",
			},
			wantErr: false,
		},
		{
			name: "failed case - noti url already exists",
			args: args{
				c: ctx,
				req: ReqGenerateAPIKey{
					CustomerID: 21314,
					Permissions: []Permission{
						{
							Name:   constant.PERMISSION_NAME,
							Action: constant.PERMISSION_ACTION,
						},
					},
				},
			},
			wantRes: ResGenerateAPIKey{
				APIKey: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := GenerateAPIKey(tt.args.c, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GenerateAPIKey() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
