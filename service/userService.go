package service

import (
	"log"

	"github.com/Yefhem/student-syllabus/model"
	"github.com/Yefhem/student-syllabus/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(uDTO model.UserDTO) error
	FindUserByEmail(email string) (model.User, error)
	VerifyPassword(hashedPassword, password string) bool
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// --------------------> Methods...
// ---------->
func (u *userService) CreateUser(uDTO model.UserDTO) error {

	user := model.User{
		Name:     uDTO.Name,
		Email:    uDTO.Email,
		Password: uDTO.Password,
	}

	if err := u.userRepo.Create(user); err != nil {
		return err
	}

	return nil
}

func (u *userService) FindUserByEmail(email string) (model.User, error) {

	user, err := u.userRepo.FindUserByEmail(email)
	if err != nil {
		return user, err
	}

	return user, nil

}

func (u *userService) VerifyPassword(hashedPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		log.Println(err)
		return false
	}
	return true
}
