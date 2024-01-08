package router

import (
	"context"

	"github.com/loak155/microservices-auth-service/usecase"
	"github.com/loak155/microservices-proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type authGRPCServer struct {
	pb.UnimplementedAuthServiceServer
	uu usecase.IAuthUsecase
}

func NewAuthGRPCServer(grpcServer *grpc.Server, uu usecase.IAuthUsecase) pb.AuthServiceServer {
	s := authGRPCServer{uu: uu}
	pb.RegisterAuthServiceServer(grpcServer, &s)
	reflection.Register(grpcServer)
	return &s
}

func (s *authGRPCServer) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	res := pb.GenerateTokenResponse{}
	authRes, err := s.uu.GenerateToken(int(req.UserId))
	if err != nil {
		return nil, err
	}
	res.Token = authRes
	return &res, nil
}

func (s *authGRPCServer) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	res := pb.ValidateTokenResponse{}
	authRes, err := s.uu.ValidateToken(req.Token)
	if err != nil {
		return nil, err
	}
	res.Valid = authRes
	return &res, nil
}

func (s *authGRPCServer) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	res := pb.RefreshTokenResponse{}
	authRes, err := s.uu.RefreshToken(req.Token)
	if err != nil {
		return nil, err
	}
	res.Token = authRes
	return &res, nil
}
