package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 1.
	// //sugar
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name::%s age::%d	", "nguyen quang", 18)
	// //logger
	// logger := zap.NewExample()
	// logger.Info("hello", zap.String("name", "thanh quang"), zap.Int("age", 19))

	// 2.
	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello Example")
	// //develop
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")
	// //prodution
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// 3. 	Custom zap
	encoder := GetEncoderLog()
	writeSync := getWriteLogsSync()
	core := zapcore.NewCore(encoder, writeSync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// format logger
func GetEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	//conver 1724381960.2213519 -> 2024-08-23T09:59:20.219+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	//change title ts -> time
	encoderConfig.TimeKey = "time"

	// change level from info -> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	//"caller":"cli/main.logs.go:21"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteLogsSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./logs/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)             //write file
	syncFileConsole := zapcore.AddSync(os.Stderr) //write console

	return zapcore.NewMultiWriteSyncer(syncFileConsole, syncFile)
}
