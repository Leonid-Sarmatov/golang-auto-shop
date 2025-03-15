package server

import (
	"context"
	"fmt"
	"golang_auto_shop/internal/core/models"
	"golang_auto_shop/internal/core/user"
	. "golang_auto_shop/internal/transport/grpc/generated"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ooo interface { // size=16 (0x10)
    // Операции с пользователями
    AddUser(context.Context, *AddUserRequest) (*Response, error)
    DeleteUser(context.Context, *IDRequest) (*Response, error)
    UpdateUser(context.Context, *UpdateUserRequest) (*Response, error)
    GetUser(context.Context, *IDRequest) (*User, error)
    // Операции с двигателями
    AddEngine(context.Context, *AddEngineRequest) (*Response, error)
    DeleteEngine(context.Context, *IDRequest) (*Response, error)
    UpdateEngine(context.Context, *UpdateEngineRequest) (*Response, error)
    GetEngine(context.Context, *IDRequest) (*Engine, error)
    // Операции с автомобилями
    AddCarModel(context.Context, *AddCarModelRequest) (*Response, error)
    DeleteCarModel(context.Context, *IDRequest) (*Response, error)
    UpdateCarModel(context.Context, *UpdateCarModelRequest) (*Response, error)
    GetCarModel(context.Context, *IDRequest) (*CarModel, error)
    // Операции связывания
    AddCarToUser(context.Context, *AddCarToUserRequest) (*Response, error)
    RemoveCarFromUser(context.Context, *RemoveCarFromUserRequest) (*Response, error)
    GetUserCars(context.Context, *GetUserCarsRequest) (*CarModels, error)
    mustEmbedUnimplementedCarShopServer()
}



type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
	userLogicCore user.UserLogicCore
	CarShopServer
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
	RegisterCarShopServer(server.grpcServer, server)

	// Старт сервера
	log.Println("Starting gRPC server on :40001")
	return server.grpcServer.Serve(server.listener)
}

func (server *Server) Stop() {
	log.Println("Stopping gRPC server...")
	server.grpcServer.GracefulStop()
}

/*
=======================================================
================= Обработчики для User ================
=======================================================
*/
func (server *Server)AddUser(context context.Context, request *AddUserRequest) (*Response, error) {
	err := server.userLogicCore.AddUser(models.User{
		Name: request.Name,
		Email: request.Email,
		Cars: make([]models.CarModel, 0),
		CreatedAt: time.Now(),
	})

	if err != nil {
		return &Response{
			Success: true,
			Message: "User creating was succussful",
		}, nil
	}

	return &Response{
		Success: false,
		Message: "User creating was failed",
	}, nil
}

func (server *Server)DeleteUser(context context.Context, request *IDRequest) (*Response, error) {
	err := server.userLogicCore.DeleteUser(request.Id)

	if err != nil {
		return &Response{
			Success: true,
			Message: "User creating was succussful",
		}, nil
	}

	return &Response{
		Success: false,
		Message: "User creating was failed",
	}, nil
}

func (server *Server)UpdateUser(context context.Context, request *UpdateUserRequest) (*Response, error) {
	switch r := request.Update.(type) {
	case *UpdateUserRequest_Email:
		err := server.userLogicCore.UpdateUserEmail(request.Id, r.Email)
		if err != nil {
			return &Response{
				Success: false,
				Message: "",
			}, nil
		}
	case *UpdateUserRequest_Name:
		err := server.userLogicCore.UpdateUserName(request.Id, r.Name)
		if err != nil {
			return &Response{
				Success: false,
				Message: "",
			}, nil
		}
	default:
		return &Response{
			Success: false,
			Message: "",
		}, nil
	}

	return &Response{
		Success: true,
		Message: "",
	}, nil
}

func (server *Server)GetUser(context context.Context, request *IDRequest) (*User, error) {
	u, err := server.userLogicCore.GetUser(request.Id)
	if err != nil {
		return nil, fmt.Errorf("")
	}

	return &User{
		Id: strconv.Itoa(int(u.ID)),
		Name: u.Name,
		Email: u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt),
	}, nil
}