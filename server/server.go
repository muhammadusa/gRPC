package main

import (

	"log"
	"net"
	// "errors"

	"app/proto"
	"app/pkg"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedMathServer
}

func (*server) Max(r *proto.Request, s proto.Math_MaxServer) error {

	s.Send(&proto.Response{ Number: pkg.Divide(r.GetNumbers(), 10, )})

	// return errors.New("Message")
	return nil
}

func main () {

	lis, err := net.Listen("tcp", ":9090")

	if err != nil { log.Fatalf("%v", err) }

	s := grpc.NewServer()

	proto.RegisterMathServer(s, &server{})

	err = s.Serve(lis)

	if err != nil { log.Fatalf("%v", err) }
}