package log

import (
	"github.com/josepmdc/goboilerplate/conf"
	"go.uber.org/zap"
)

// StandardLogger represents a common interface for logging. It has the
// the most important functions for writing logs. Most logger are
// implementations of this interface, so you can just swap them without
// much change.
type StandardLogger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

var Logger StandardLogger

// ConfigureLogger creates a new implementation of the Logger interface
// based on the configuration
func ConfigureLogger(cfg *conf.LogConfig) error {
	zapCfg := zap.NewProductionConfig()
	if cfg.Debug {
		zapCfg = zap.NewDevelopmentConfig()
	}
	zapCfg.OutputPaths = append(zapCfg.OutputPaths, cfg.File)
	logger, err := zapCfg.Build()
	if err != nil {
		return err
	}
	Logger = logger.Sugar()
	return nil
}
