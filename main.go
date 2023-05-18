package main

import (
	"fmt"
	"log"
	"net"

	protos "github.com/kamalbahsine/simple-grpc-app/protos/podinfos"
	"github.com/kamalbahsine/simple-grpc-app/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
        
	func NewPodInfo() *PodInfo {
	return &PodInfo{}
         }  



	// create new gRPC server
	s := grpc.NewServer()

	// create new instance of PodInfo server
	pinf := server.NewPodInfo()
        
	// register reflection API https://github.com/grpc/grpc/blob/master/doc/server-reflection.md
	reflection.Register(s)

	// register it to the grpc server
	protos.RegisterPodInfoServer(s, pinf)

	// create socket to listen to requests
	tl, err := net.Listen("tcp", "0.0.0.0:8765")
	if err != nil {
		log.Fatal(fmt.Println("Error starting tcp listener on port 8765", err))
	}

	// start listening
	s.Serve(tl)
}
