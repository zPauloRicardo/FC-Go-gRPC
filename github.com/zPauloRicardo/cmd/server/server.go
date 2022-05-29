package main

import (
	"gRPC/github.com/zPauloRicardo/pb"
	"gRPC/github.com/zPauloRicardo/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	//implementa um listener para a porta
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Não foi possivel conectar: %v", err)
	}

	//cria um server grpc
	grpcServer := grpc.NewServer()

	//registra o service no server grcp
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	//atribui o listener de porta ao servidor grcp e inicia o mesmo
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Não foi possivel servir: %v", err)
	}
}
