package services

import (
	"testing"

	"pruse_logs/log-wrapper/internal/fixtures/test_helpers"

	"github.com/FenixAra/go-util/testh"
)

func TestPing(t *testing.T) {
	p := NewPing(test_helpers.TestInit())

	err := p.Ping()
	testh.AssertNoErr("Unable to ping database", err, t)
}
