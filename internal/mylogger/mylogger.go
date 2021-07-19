package mylogger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewLogger(devMode bool) (*zap.SugaredLogger, error) {
	var logger *zap.Logger
	var err error
	if devMode {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, errors.Wrap(err, "failed to create development logger")
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, errors.Wrap(err, "failed to create production logger")
		}
	}
	defer logger.Sync()
	return logger.Sugar(), nil
}
