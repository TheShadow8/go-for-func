package services

import (
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/asaskevich/govalidator.v9"

	"gitlab.com/TheShadow8/go-test-fiber/models"
	"gitlab.com/TheShadow8/go-test-fiber/repository"
	"gitlab.com/TheShadow8/go-test-fiber/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthServices interface {
	SignUp(user *models.User) error
	SignIn(input *models.User) (user *models.User, token string, error error)
	GetByEmail(email string) (*models.User, error)
	GetUser(userId string) (*models.User, error)
	GetUsers() ([]*models.User, error)
}

type authServices struct {
	usersRepo repository.UsersRepository
}

func NewAuthServices(usersRepo repository.UsersRepository) AuthServices {
	return &authServices{usersRepo}
}

func (s *authServices) SignUp(user *models.User) error {
	var err error

	user.Email = util.NormalizeEmail(user.Email)

	if !govalidator.IsEmail(user.Email) {
		return util.ErrInvalidEmail
	}

	exitedUser, err := s.GetByEmail(user.Email)

	if err != nil {
		return util.ErrInvalidInput
	}

	if exitedUser != nil {
		return util.ErrEmailAlreadyExists
	}

	if strings.TrimSpace(user.Password) == "" {
		return util.ErrEmptyPassword
	}

	user.Password, err = util.EncryptPassword(user.Password)

	if err != nil {
		return err
	}

	now := time.Now()

	user.ID = primitive.NewObjectID()
	user.CreatedAt = &now
	user.UpdatedAt = user.CreatedAt

	_, err = s.usersRepo.Save(user)

	if err != nil {
		return err
	}

	return nil
}

func (s *authServices) SignIn(input *models.User) (user *models.User, token string, error error) {

	input.Email = util.NormalizeEmail(input.Email)
	user, err := s.GetByEmail(input.Email)

	if err != nil {
		log.Printf("%s signin verify password failed: %v\n", input.Email, err.Error())
		return nil, "", util.ErrInvalidCredentials
	}

	if user == nil {
		log.Printf("%s signin verify password failed: \n", input.Email)

		return nil, "", util.ErrInvalidCredentials
	}

	err = util.VerifyPassword(user.Password, input.Password)

	if err != nil {
		log.Printf("%s signin verify password failed: %v\n", input.Email, err.Error())
		return nil, "", util.NewAppError(util.ErrInvalidCredentials, http.StatusUnauthorized)
	}

	token, err = util.NewToken(user.ID.String())

	if err != nil {
		log.Printf("%s signin create token failed: %v\n", input.Email, err.Error())
		return nil, "", util.NewAppError(err, http.StatusUnauthorized)
	}

	return user.SanitizeUser(), "Bearer " + token, nil
}

func (s *authServices) GetUser(userId string) (*models.User, error) {
	return s.usersRepo.GetById(userId)
}

func (s *authServices) GetByEmail(email string) (*models.User, error) {
	return s.usersRepo.GetByEmail(email)
}

func (s *authServices) GetUsers() ([]*models.User, error) {
	return s.usersRepo.GetAll()
}
