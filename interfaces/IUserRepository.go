package interfaces

import "github.com/barretodotcom/go-grpc-integration/entities"

type IUserRepository interface {
	CreateUser(*entities.User) error
	CountUserByEmail(string) (int64, error)
	FindUserByEmail(email string) (*entities.User, error)
}
