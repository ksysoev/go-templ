package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {
	chats := NewMockchatRepo(t)
	someAPI := NewMocksomeAPIProv(t)
	svc := New(chats, someAPI)

	assert.NotNil(t, svc, "New() should return a non-nil Service instance")
}

func TestService_CheckHealth(t *testing.T) {
	tests := []struct {
		setupMocks func(t *testing.T, chats *MockchatRepo, someAPI *MocksomeAPIProv)
		name       string
		wantErr    bool
	}{
		{
			name: "Success",
			setupMocks: func(t *testing.T, chats *MockchatRepo, someAPI *MocksomeAPIProv) {
				t.Helper()
				chats.EXPECT().CheckHealth(mock.Anything).Return(nil)
				someAPI.EXPECT().CheckHealth(mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "chatRepo failure",
			setupMocks: func(t *testing.T, chats *MockchatRepo, someAPI *MocksomeAPIProv) {
				t.Helper()
				chats.EXPECT().CheckHealth(mock.Anything).Return(assert.AnError)
				someAPI.EXPECT().CheckHealth(mock.Anything).Return(nil)
			},
			wantErr: true,
		},
		{
			name: "someAPI failure",
			setupMocks: func(t *testing.T, chats *MockchatRepo, someAPI *MocksomeAPIProv) {
				t.Helper()
				chats.EXPECT().CheckHealth(mock.Anything).Return(nil)
				someAPI.EXPECT().CheckHealth(mock.Anything).Return(assert.AnError)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chats := NewMockchatRepo(t)
			someAPI := NewMocksomeAPIProv(t)
			s := New(chats, someAPI)

			tt.setupMocks(t, chats, someAPI)

			err := s.CheckHealth(t.Context())

			if tt.wantErr {
				assert.Error(t, err, "CheckHealth() should return an error")
			} else {
				assert.NoError(t, err, "CheckHealth() should not return an error")
			}
		})
	}
}
