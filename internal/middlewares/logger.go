package middlewares

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"os"
	"time"
)

const (
	LoggerKey  string = "logger"
	TraceIdKey string = "traceId"
	SpanKey    string = "span"
)

var (
	Logger *zap.Logger
)

func InitLogger() func() error {
	config := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	config.EncodeDuration = zapcore.MillisDurationEncoder
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(os.Stdout), zapcore.InfoLevel)
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return nil
}

func GrpcAccessLog() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		begin := time.Now()
		defer func() {
			CtxInfo(ctx, "access request", zap.Reflect("req", req), zap.Reflect("res", resp),
				zap.String("method", info.FullMethod), zap.Error(err), zap.Duration("cost", time.Since(begin)),
			)
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}

func GrpcSetContextLogger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		value := ctx.Value(TraceIdKey)
		if value != nil {
			if traceId, ok := value.(string); ok {
				loggerInstance := Logger.With(
					zap.String("traceId", traceId),
				)
				ctx = context.WithValue(ctx, LoggerKey, loggerInstance)
			}
		}
		return handler(ctx, req)
	}
}

func CtxInfo(ctx context.Context, msg string, fields ...zap.Field) {
	value := ctx.Value(LoggerKey)
	if loggerInstantce, ok := value.(*zap.Logger); ok {
		loggerInstantce.Info(msg, fields...)
		return
	}
	Logger.Info(msg, fields...)
}
func CtxWarn(ctx context.Context, msg string, fields ...zap.Field) {
	value := ctx.Value(LoggerKey)
	if loggerInstantce, ok := value.(*zap.Logger); ok {
		loggerInstantce.Warn(msg, fields...)
		return
	}
	Logger.Warn(msg, fields...)
}

func CtxError(ctx context.Context, msg string, fields ...zap.Field) {
	value := ctx.Value(LoggerKey)
	if loggerInstantce, ok := value.(*zap.Logger); ok {
		loggerInstantce.Error(msg, fields...)
		return
	}
	Logger.Error(msg, fields...)
}
