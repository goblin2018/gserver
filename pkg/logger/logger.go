package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

/**
Dev:
	zapcore.CapitalColorLevelEncoder

调试时，使用彩色打印
正常运行时，指定日志存储位置
*/

func init() {
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "test.logger",
		MaxSize:    500, // megabytes default 100
		MaxAge:     0,   // days
		MaxBackups: 3,
	})

	ecfg := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "line",
		StacktraceKey: "trace",
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,

		//EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			type appendTimeEncoder interface {
				AppendTimeLayout(time.Time, string)
			}
			layout := "2006-01-02 15:04:05.000"
			if enc, ok := enc.(appendTimeEncoder); ok {
				enc.AppendTimeLayout(t, layout)
				return
			}
			enc.AppendString(t.Format(layout))
		},
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     nil,
	}

	var syncer zapcore.WriteSyncer
	// 使用多个writer
	syncer = zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(writer),
	)
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(ecfg), syncer, zap.DebugLevel)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()

}
