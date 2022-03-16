package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func New(structuredLogging bool, projectName string) (*zap.SugaredLogger, error) {
	var logger *zap.Logger
	var err error
	if structuredLogging {
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, errors.Wrap(err, "failed to create production logger")
		}
	} else {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, errors.Wrap(err, "failed to create development logger")
		}
	}
	defer logger.Sync()

	return logger.Sugar(), nil
}
