package chat

import (
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {
	daoMock := NewMockchatDAO(t)
	repo := New(daoMock)

	assert.NotNil(t, repo, "New() should return a non-nil ChatRepo instance")
}

func TestChatRepo_CheckHealth(t *testing.T) {
	tests := []struct {
		setupMocks func(t *testing.T, dao *MockchatDAO)
		name       string
		wantErr    bool
	}{
		{
			name: "success",
			setupMocks: func(t *testing.T, dao *MockchatDAO) {
				t.Helper()
				dao.EXPECT().Ping(mock.Anything).Return(redis.NewStatusResult("", nil))
			},
			wantErr: false,
		},
		{
			name: "fail",
			setupMocks: func(t *testing.T, dao *MockchatDAO) {
				t.Helper()
				dao.EXPECT().Ping(mock.Anything).Return(redis.NewStatusResult("", assert.AnError))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			daoMock := NewMockchatDAO(t)
			c := New(daoMock)

			tt.setupMocks(t, daoMock)

			err := c.CheckHealth(t.Context())

			if tt.wantErr {
				assert.Error(t, err, "CheckHealth() should return an error")
			} else {
				assert.NoError(t, err, "CheckHealth() should not return an error")
			}
		})
	}
}
