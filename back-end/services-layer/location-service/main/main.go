package main

import (
	"location/myInit"
	"location/mylog"
	"location/proto"
	"location/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	mylog.LogClient.Infof("Location service start..")

	defer myInit.Db.Close()

	//Listen port
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		mylog.LogClient.Errorf("| net.listen | %v |", err)
	}

	//Init gPRC server
	grpcServer := grpc.NewServer()
	mylog.LogClient.Infof("GrpcServer have been create")

	mylog.LogClient.Infof("Now start service")
	myServer := service.LocationService{}

	//Register location service
	proto.RegisterLocationServiceServer(grpcServer, &myServer)

	//Start
	err = grpcServer.Serve(lis)
	if err != nil {
		mylog.LogClient.Errorf("| net.listen | %v |", err)
	}
}
