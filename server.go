package main

import (

	"github.com/JIeeiroSst/app/proto"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"context"
)

type Server struct {}

func main(){
	lis,err:=net.Listen("tcp",":8080")
	if err!=nil {
		log.Println(err)
	}
	grpServer:=grpc.NewServer()
	proto.RegisterAddServiceServer(grpServer,&Server{})
	reflection.Register(grpServer)
	if err:=grpServer.Serve(lis);err!=nil{
		log.Println(err)
	}
}

func (s *Server) Add(ctx context.Context,request *proto.Request) (*proto.Response,error){
	a,b:=request.GetA(),request.GetB()
	result:=a+b
	return &proto.Response{Result:result},nil
}

func (s *Server) Multiply(ctx context.Context,request *proto.Request) (*proto.Response,error){
	a,b:=request.GetA(),request.GetB()
	result:=a+b
	return &proto.Response{Result:result},nil
}

