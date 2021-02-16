package services

import (
	"strings"
	"time"

	"gopkg.in/asaskevich/govalidator.v9"

	"gitlab.com/TheShadow8/go-test-fiber/models"
	"gitlab.com/TheShadow8/go-test-fiber/repository"
	"gitlab.com/TheShadow8/go-test-fiber/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthServices interface {
	SignUp(user *models.User) (*models.User, error)
	GetUser(userId string) (*models.User, error)
}

type authServices struct {
	usersRepo repository.UsersRepository
}

func NewAuthServices(usersRepo repository.UsersRepository) AuthServices {
	return &authServices{usersRepo}
}

func (s *authServices) SignUp(user *models.User) (*models.User, error) {
	var err error

	user.Email = util.NormalizeEmail(user.Email)

	if !govalidator.IsEmail(user.Email) {
		return nil, util.ErrInvalidEmail
	}

	if strings.TrimSpace(user.Password) == "" {
		return nil, util.ErrEmptyPassword
	}

	user.Password, err = util.EncryptPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	inserted, err := s.usersRepo.Save(user)
	if err != nil {
		return nil, err
	}

	if oid, ok := inserted.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}

	return user.SanitizeUser(), nil
}

func (s *authServices) GetUser(userId string) (*models.User, error) {
	return s.usersRepo.GetById(userId)
}
