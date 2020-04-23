package logic

import (
	"context" 
	"kfd/pkg/gerrors" 
	"fmt"
	"kfd/pkg/util" 
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func logPanic(serverName string, ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, err *error) {
	p := recover()
	if p != nil {
		//logger.Logger.Error(serverName+" panic", zap.Any("info", info), zap.Any("ctx", ctx), zap.Any("req", req), zap.Any("panic", p), zap.String("stack", util.GetStackInfo()))

		fmt.Println(serverName+" panic")
		fmt.Println(info)
		fmt.Println(ctx)
		fmt.Println(req)
		fmt.Println(p) 
		fmt.Println(util.GetStackInfo())
 
		*err = gerrors.ErrUnknown
	}
}

// 服务器端的单向调用的拦截器
func LogicIntInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		logPanic("logic_int_interceptor", ctx, req, info, &err)
	}()

	resp, err = handler(ctx, req)
 
	s, _ := status.FromError(err)
	if s.Code() != 0 && s.Code() < 1000 {
		md, _ := metadata.FromIncomingContext(ctx)

		//logger.Logger.Error("logic_int_interceptor", zap.String("method", info.FullMethod), zap.Any("md", md), zap.Any("req", req), zap.Any("resp", resp), zap.Error(err), zap.String("stack", gerrors.GetErrorStack(s)))
		fmt.Println( " logic_int_interceptor")
		fmt.Println("method", info.FullMethod)
		fmt.Println(md)
		fmt.Println(req)
		fmt.Println(resp) 
		fmt.Println(err) 
		fmt.Println(gerrors.GetErrorStack(s))
	}
	return resp, err
}
 