package ulog

import (
	"os"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

// NewZap ...
func NewZap() (*zap.Logger, error) {
	if os.Getenv("GO_ENV") == "development" || os.Getenv("GO_ENV") == "test" {
		return zap.NewDevelopment()
	}
	projectId := os.Getenv("GOOGLE_PROJECT_ID")
	return zapdriver.NewProductionWithCore(zapdriver.WrapCore(
		zapdriver.ReportAllErrors(true),
		zapdriver.ServiceName(projectId),
	))
}
