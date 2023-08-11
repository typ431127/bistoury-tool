package core

import "go.uber.org/zap"

func Zap() (sugar *zap.SugaredLogger) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar = logger.Sugar()
	return sugar
}
