package logger

import (
	"context"

	"github.com/blendle/zapdriver"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
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

	ctx := context.Background()
	sc := trace.SpanContextFromContext(ctx)
	fields := zapdriver.TraceContext(sc.TraceID().String(), sc.SpanID().String(), sc.IsSampled(), projectName)
	logger = logger.With(fields...)
	return logger.Sugar(), nil
}
