package provider

import (
	"context"

	"github.com/barretodotcom/go-grpc-integration/entities"
	"github.com/barretodotcom/go-grpc-integration/pb"
	"github.com/barretodotcom/go-grpc-integration/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func (s *UserServer) RegisterUser(ctx context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {

	userRepository := repositories.NewUserRepository()

	usersWithThisEmail, err := userRepository.CountUserByEmail(r.GetEmail())
	if err != nil {
		return &pb.UserResponse{
			Sucess:  false,
			Message: err.Error(),
		}, nil
	}

	if usersWithThisEmail != 0 {
		return &pb.UserResponse{
			Sucess:  false,
			Message: "One user is already registered with this email.",
		}, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.MinCost)

	if err != nil {
		return &pb.UserResponse{
			Sucess:  false,
			Message: err.Error(),
		}, nil
	}

	user := &entities.User{
		Name:     r.GetName(),
		Email:    r.GetEmail(),
		Password: string(hashedPassword),
	}

	err = userRepository.CreateUser(user)

	if err != nil {
		return &pb.UserResponse{
			Sucess:  false,
			Message: err.Error(),
		}, nil
	}

	return &pb.UserResponse{
		Sucess: true,
	}, nil
}
