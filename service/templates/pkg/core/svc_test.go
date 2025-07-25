package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {
	users := NewMockuserRepo(t)
	someAPI := NewMocksomeAPIProv(t)

	svc := New(users, someAPI)

	assert.NotNil(t, svc, "New() should return a non-nil Service instance")
}

func TestService_CheckHealth(t *testing.T) {
	tests := []struct {
		name       string
		setupMocks func(t *testing.T, users *MockuserRepo, someAPI *MocksomeAPIProv)
		wantErr    bool
	}{
		{
			name: "Success",
			setupMocks: func(t *testing.T, users *MockuserRepo, someAPI *MocksomeAPIProv) {
				users.EXPECT().CheckHealth(mock.Anything).Return(nil)
				someAPI.EXPECT().CheckHealth(mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "userRepo failure",
			setupMocks: func(t *testing.T, users *MockuserRepo, someAPI *MocksomeAPIProv) {
				users.EXPECT().CheckHealth(mock.Anything).Return(assert.AnError)
				someAPI.EXPECT().CheckHealth(mock.Anything).Return(nil)
			},
			wantErr: true,
		},
		{
			name: "someAPI failure",
			setupMocks: func(t *testing.T, users *MockuserRepo, someAPI *MocksomeAPIProv) {
				users.EXPECT().CheckHealth(mock.Anything).Return(nil)
				someAPI.EXPECT().CheckHealth(mock.Anything).Return(assert.AnError)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := NewMockuserRepo(t)
			someAPI := NewMocksomeAPIProv(t)

			tt.setupMocks(t, users, someAPI)

			s := New(users, someAPI)
			err := s.CheckHealth(t.Context())
			if tt.wantErr {
				assert.Error(t, err, "CheckHealth() should return an error")
			} else {
				assert.NoError(t, err, "CheckHealth() should not return an error")
			}
		})
	}
}
