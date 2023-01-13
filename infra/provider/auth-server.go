package provider

import (
	"context"

	"github.com/barretodotcom/go-grpc-integration/pb"
	"github.com/barretodotcom/go-grpc-integration/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
}

func (a *AuthServer) AuthUser(ctx context.Context, r *pb.AuthRequest) (*pb.AuthResponse, error) {

	userRepository := repositories.NewUserRepository()

	userEmail := r.GetEmail()

	user, err := userRepository.FindUserByEmail(userEmail)

	if user == nil {
		return &pb.AuthResponse{
			Token:   "",
			Sucess:  false,
			Message: "User not found.",
		}, nil
	}

	if err != nil {
		return &pb.AuthResponse{
			Token:   "",
			Sucess:  false,
			Message: err.Error(),
		}, nil
	}

	password := r.GetPassword()

	sucess := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if sucess != nil {
		return &pb.AuthResponse{
			Token:   "",
			Sucess:  false,
			Message: "Wrong password",
		}, nil
	}

	return &pb.AuthResponse{
		Token:   "",
		Sucess:  true,
		Message: "Authorized.",
	}, nil
}
