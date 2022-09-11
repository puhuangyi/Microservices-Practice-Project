package main

import (
	"login/myInit"
	"login/mylog"
	"login/proto"
	"login/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	mylog.LogClient.Infof("Login service start..")

	defer myInit.Db.Close()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		mylog.LogClient.Errorf("| net.listen | %v |", err)
	}

	grpcServer := grpc.NewServer()
	mylog.LogClient.Infof("GrpcServer have been create")

	mylog.LogClient.Infof("Now start service")
	myServer := service.LoginClient{}
	proto.RegisterLoginServiceServer(grpcServer, &myServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		mylog.LogClient.Errorf("| net.listen | %v |", err)
	}
}
