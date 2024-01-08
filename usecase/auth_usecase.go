package usecase

import (
	"fmt"
	"os"

	"github.com/loak155/microservices-auth-service/client"
	"github.com/loak155/microservices-auth-service/utils"
	"github.com/loak155/microservices-proto/pb"
)

type IAuthUsecase interface {
	GenerateToken(user_id int) (string, error)
	ValidateToken(token string) (bool, error)
	RefreshToken(token string) (string, error)
}

type authUsecase struct {
	uc         client.IUserGRPCClient
	jwtManager *utils.JwtManager
}

func NewAuthUsecase(uc client.IUserGRPCClient, jwtManager utils.JwtManager) IAuthUsecase {
	return &authUsecase{uc, jwtManager}
}

func (uu *authUsecase) GenerateToken(user_id int) (string, error) {
	req := pb.GetUserRequest{Id: int32(user_id)}
	res, err := uu.uc.GetUser(&req)
	if err != nil {
		return "", err
	}
	// TODO: resのパスワードが一致するか確認する
	fmt.Println(res)
	token, err := uu.jwtManager.Generate(user_id)
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uu *authUsecase) ValidateToken(token string) (bool, error) {
	_, err := uu.jwtManager.ValidateToken(token)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (uu *authUsecase) RefreshToken(token string) (string, error) {
	claims, err := uu.jwtManager.ValidateToken(token)
	if err != nil {
		return "", err
	}
	refreshToken, err := uu.jwtManager.Generate(claims.UserId)
	if err != nil {
		return "", err
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return refreshTokenString, nil
}
