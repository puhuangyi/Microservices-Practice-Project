package main

import (
	"net"
	"payment/myServices"
	"payment/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	defer myServices.Db.Close()

	myServer := myServices.MyService{}

	grpcServer := grpc.NewServer()

	proto.RegisterPaymentServiceServer(grpcServer, &myServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

}
