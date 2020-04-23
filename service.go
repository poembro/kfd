package main

import (
	logic "kfd/api/logic/grpc"
	"kfd/internal/logic/db"
	"net"
	"google.golang.org/grpc"
	"kfd/pkg/util"
	"kfd/pkg/pb"
)

// StartRpcServer 启动rpc服务
func StartRpcServer() {
	go func() {
		defer util.RecoverPanic()

		intListen, err := net.Listen("tcp", ":50000")
		if err != nil {
			panic(err)
		}
		s := grpc.NewServer(grpc.UnaryInterceptor(logic.LogicIntInterceptor))
		pb.RegisterLogicUserServer(s, &logic.LogicUserServer{})
		err = s.Serve(intListen)
		if err != nil {
			panic(err)
		}
	}()
}

func main() {
	// 初始化数据库
	db.InitDB()

	// 初始化自增id配置
	util.InitUID(db.DBCli)


	StartRpcServer()
	
	select {}
}




