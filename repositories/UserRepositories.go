package repositories

import (
	"context"
	"errors"

	"github.com/barretodotcom/go-grpc-integration/entities"
	"github.com/barretodotcom/go-grpc-integration/infra/db"
	"github.com/barretodotcom/go-grpc-integration/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	interfaces.IUserRepository
	UsersDB mongo.Collection
}

func NewUserRepository() interfaces.IUserRepository {
	return &UserRepository{
		UsersDB: *db.DB,
	}
}

func (repo *UserRepository) CreateUser(user *entities.User) error {
	_, err := repo.UsersDB.InsertOne(context.TODO(), &user)

	return err
}

func (repo *UserRepository) CountUserByEmail(email string) (int64, error) {
	count, err := repo.UsersDB.CountDocuments(context.TODO(), bson.D{{Key: "email", Value: email}})

	return count, err
}

func (repo *UserRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user *entities.User
	singleResult := repo.UsersDB.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}})

	if errors.Is(singleResult.Err(), mongo.ErrNoDocuments) {
		return nil, nil
	}

	err := singleResult.Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
