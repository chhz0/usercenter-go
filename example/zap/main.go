package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Infow("print zap info lv message", "url", "go.uber.org/zap")
	sugar.Warnw("print zap warn lv message", "url", "go.uber.org/zap")
	sugar.Errorw("print zap error lv message", "url", "go.uber.org/zap")
}
