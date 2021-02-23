package repository

import (
	"context"
	"errors"

	"gitlab.com/TheShadow8/go-test-fiber/db"
	"gitlab.com/TheShadow8/go-test-fiber/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UsersCollection = "users"

type UsersRepository interface {
	Save(user *models.User) (*mongo.InsertOneResult, error)
	// Update(user *models.User) error
	GetById(id string) (user *models.User, error error)
	GetByEmail(email string) (user *models.User, err error)
	// GetAll() (users []*models.User, err error)
	// Delete(id string) error
}

type usersRepository struct {
	c *mongo.Collection
}

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{conn.DB().Collection(UsersCollection)}
}

func (r *usersRepository) Save(user *models.User) (*mongo.InsertOneResult, error) {

	return r.c.InsertOne(context.TODO(), user)
}

// func (r *usersRepository) Update(user *models.User) error {
// 	return r.c.UpdateId(user.Id, user)
// }

func (r *usersRepository) GetById(id string) (user *models.User, error error) {
	userID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		return nil, err
	}
	documentReturned := r.c.FindOne(context.TODO(), bson.M{"_id": userID})

	userDecode := models.User{}

	err = documentReturned.Decode(&userDecode)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Not Found")
		}

		return nil, err
	}

	return userDecode.SanitizeUser(), nil
}

func (r *usersRepository) GetByEmail(email string) (user *models.User, err error) {
	documentReturned := r.c.FindOne(context.TODO(), bson.M{"email": email})

	userDecode := models.User{}

	err = documentReturned.Decode(&userDecode)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Not Found")
		}
		return nil, err
	}

	return & userDecode, nil
}

// func (r *usersRepository) GetAll() (users []*models.User, err error) {
// 	err = r.c.Find(bson.M{}).All(&users)
// 	return users, err
// }

// func (r *usersRepository) Delete(id string) error {
// 	return r.c.RemoveId(bson.ObjectIdHex(id))
// }
