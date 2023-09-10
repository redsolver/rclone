// Test Sia filesystem interface
package siad_test

import (
	"testing"

	"github.com/rclone/rclone/backend/siad"

	"github.com/rclone/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSiad:",
		NilObject:  (*sia.Object)(nil),
	})
}
