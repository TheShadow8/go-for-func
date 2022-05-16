package repository

import (
	"context"
	"fmt"

	"gitlab.com/TheShadow8/go-test-fiber/db"
	"gitlab.com/TheShadow8/go-test-fiber/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type AbcRepository[R  models.File | models.User]interface {
	Save(*R) (*mongo.InsertOneResult, error)
	GetById(id string) (r *R, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (users []*R, err error)
	// Update(user *models.User) error
	// Delete(id string) error
}

type abcRepository[T models.File | models.User] struct {
	c *mongo.Collection
}

func NewAbcRepository[Y models.File | models.User](conn db.Connection, collName string) AbcRepository[Y] {
	return &abcRepository[Y]{conn.DB().Collection(collName)}
}

func (r *abcRepository[T]) Save(user *T) (*mongo.InsertOneResult, error) {

	return r.c.InsertOne(context.TODO(), user)
}

// func (r *usersRepository) Update(user *models.User) error {
// 	return r.c.UpdateId(user.Id, user)
// }

func (r *abcRepository[T]) GetById(id string) (abc *T, error error) {
	userID, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		return nil, err
	}
	documentReturned := r.c.FindOne(context.TODO(), bson.M{"_id": userID})

	var userDecode T

	fmt.Println("DR", documentReturned)

	err = documentReturned.Decode(&userDecode)

	fmt.Println("Us", userDecode)


	if err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &userDecode, nil
}

func (r *abcRepository[T]) GetByEmail(email string) (user *models.User, err error) {
	documentReturned := r.c.FindOne(context.TODO(), bson.M{"email": email})

	userDecode := models.User{}

	err = documentReturned.Decode(&userDecode)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &userDecode, nil
}

func (r *abcRepository[T]) GetAll() (users []*T, err error) {
	cursor, err := r.c.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	if cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil

}
