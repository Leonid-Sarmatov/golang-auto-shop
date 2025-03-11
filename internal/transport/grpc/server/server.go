package server

import (
	"golang_auto_shop/internal/transport/grpc/generated"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
	generated.CarShopServer
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Start() error {
	// Создание слушателя для порта
	lis, err := net.Listen("tcp", ":40001")
	if err != nil {
		log.Printf("Can not open tcp port %v", err)
		return err
	}

	// Инициализация полей и регистрация сервера
	server.listener = lis
	server.grpcServer = grpc.NewServer()
	generated.RegisterCarShopServer(server.grpcServer, server)

	// Старт сервера
	log.Println("Starting gRPC server on :40001")
	return server.grpcServer.Serve(server.listener)
}

func (server *Server) Stop() {
	log.Println("Stopping gRPC server...")
	server.grpcServer.GracefulStop()
}