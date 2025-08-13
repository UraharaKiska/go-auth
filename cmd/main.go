package main

import (
	"context"
	"flag"
	"os"

	"log"

	"github.com/UraharaKiska/go-auth/internal/app"
	"github.com/UraharaKiska/go-auth/internal/logger"
	"github.com/UraharaKiska/go-auth/internal/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const serviceName = "go-auth"


func main() {
	ctx := context.Background()
	configPath := flag.String("config-path", "local.env", "path to config file")
    logLevel := flag.String("l", "info", "logging level")
    flag.Parse()
	app.InitConfigPath(*configPath)
	logger.Init(getCore(getAtomicLevel(logLevel)))
	shutdown := tracing.Init(serviceName)
	defer func() {
        if err := shutdown(ctx); err != nil {
            log.Fatalf("Error shutting down tracer: %v", err)
        }
    }()	
	a, err := app.NewApp(ctx)
	// log.Printf("App :%v", a)
	if err != nil {
		log.Fatalf("failed to init app%v: ", err)
	}
	err = a.Run()
	if err != nil {
		log.Fatalf("failed to init app%v: ", err)
	}
}



func getCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	// file := zapcore.AddSync(&lumberjack.Logger{
	// 	Filename:   "logs/app.log",
	// 	MaxSize:    10, // megabytes
	// 	MaxBackups: 3,
	// 	MaxAge:     7, // days
	// })

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	// fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	return zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		// zapcore.NewCore(fileEncoder, file, level),
	)
}

func getAtomicLevel(logLevel *string) zap.AtomicLevel {
	var level zapcore.Level
	if err := level.Set(*logLevel); err != nil {
		log.Fatalf("failed to set log level: %v", err)
	}

	return zap.NewAtomicLevelAt(level)
}