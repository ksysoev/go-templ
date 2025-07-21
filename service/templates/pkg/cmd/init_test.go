package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitCommand(t *testing.T) {
	cmd := InitCommand(BuildInfo{AppName: "app"})

	assert.Equal(t, "app", cmd.Use)
	assert.Contains(t, cmd.Short, "")
	assert.Contains(t, cmd.Long, "")

	require.Len(t, cmd.Commands(), 1)
	assert.Equal(t, "server", cmd.Commands()[0].Use)
}
