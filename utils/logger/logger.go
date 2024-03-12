package logger

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"shop-backend/config"
	"time"
)

var ProdL *zap.Logger
var L *zap.SugaredLogger

func init() {
	ProdL, _ = zap.NewProduction()
	if config.IsRelease() {
		L = ProdL.Sugar()
		return
	}

	zc := zap.NewDevelopmentEncoderConfig()
	zc.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zc.EncodeTime = zapcore.TimeEncoderOfLayout(fmt.Sprintf("[%s]", time.RFC3339))
	L = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zc),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	)).Sugar()
}
