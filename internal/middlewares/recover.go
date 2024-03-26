package middlewares

import (
	"attachment-service/internal/commons"
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"runtime/debug"
)

func GrpcRecover() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if e := doRecover(ctx); e != nil {
				err = commons.ServerInternalError.Withf("%v", e).GrpcErr()
			}
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}

func doRecover(ctx context.Context) interface{} {
	if e := recover(); e != nil {
		stack := string(debug.Stack())
		CtxError(ctx, "server panic", zap.Any("panicErr", e), zap.Stack("stack"))
		CtxError(ctx, stack)
		fmt.Println(stack)
		return e
	}
	return nil
}
