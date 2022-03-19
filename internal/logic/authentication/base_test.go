package authentication

import (
	"context"
	"testing"

	testutil "github.com/liemeilla/Notification_Service/test-util"
)

func init() {
	testutil.InitDBAndExternalStub()
}

func TestCheckAPIKey(t *testing.T) {
	ctx := testutil.InitContextForTest()

	type args struct {
		c      context.Context
		apiKey string
	}
	tests := []struct {
		name      string
		args      args
		wantValid bool
		wantErr   bool
	}{
		{
			name: "success case",
			args: args{
				c:      ctx,
				apiKey: "eyJjdXN0b21lcl9pZCI6MjEzMTQsInBlcm1pc3Npb25zIjpbeyJuYW1lIjoibm90aWZpY2F0aW9uIiwiYWN0aW9uIjoid3JpdGUifV19",
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "failed case - permission not valid",
			args: args{
				c:      ctx,
				apiKey: "ewoJImN1c3RvbWVyX2lkIjogMjEzMTQsCgkicGVybWlzc2lvbnMiOiBbewoJCSJuYW1lIjogInBheW1lbnQiLAoJCSJhY3Rpb24iOiAid3JpdGUiCgl9XQp9",
			},
			wantValid: false,
			wantErr:   false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			gotValid, err := CheckAPIKey(tt.args.c, tt.args.apiKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValid != tt.wantValid {
				t.Errorf("CheckAPIKey() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}
