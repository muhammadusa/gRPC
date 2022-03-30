package main

import (
	// "os"
	// "strconv"
	"log"
	"context"

	"app/proto"
	"app/pkg"

	"google.golang.org/grpc"
)

func main () {

	conn, err := grpc.Dial(":9090", grpc.WithInsecure())

	if err != nil { log.Fatalf("%v", err) }

	defer conn.Close()

	stub := proto.NewMathClient(conn)

	// limit, _ := strconv.Atoi(os.Args[1])

	req := proto.Request{ Numbers: pkg.Fill(9) }

	stream, err := stub.Max(context.Background(), &req)

	if err != nil { log.Fatalf("%v", err) }

	res, err := stream.Recv()

	if err != nil { log.Fatalf("%v", err) }

	log.Println(res.Number)
}