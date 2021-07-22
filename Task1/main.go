package main

import(
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func initZapLog() *zap.Logger {
    config := zap.NewDevelopmentConfig()
    config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    config.EncoderConfig.TimeKey = "timestamp"
    config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    logger, _ := config.Build()
    return logger
}

func main() {

	loggerMgr := initZapLog()
    zap.ReplaceGlobals(loggerMgr)
    defer loggerMgr.Sync() // flushes buffer, if any
    zap.S().Debug("START!")

	bst := Bst{}

	bst.Insert(10)
	bst.Insert(15)
	bst.Insert(6)
	bst.Insert(8)
	bst.PreOrderTraversal()

}