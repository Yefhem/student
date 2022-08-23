package repository

import (
	"log"

	"github.com/Yefhem/student-syllabus/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user model.User) error
	FindUserByEmail(email string) (model.User, error)
	CheckEmailPass(email, pass  interface{}) (model.User, error)
}

type userConnection struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		db: db,
	}
}

// --------------------> Methods...

// ----------> Creates a User and returns it if there is an error...
func (c *userConnection) Create(user model.User) error {
	user.Password = PasswordHasher(user.Password)

	if err := c.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// --------> Find a User by Email
func (c *userConnection) FindUserByEmail(email string) (model.User, error) {

	var user = model.User{}

	if err := c.db.Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (c *userConnection) CheckEmailPass(email, pass interface{}) (model.User, error) {
	var user = model.User{}

	if err := c.db.Where("email = ? AND pass = ?", email, pass).Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// --------> Hash to User Pass...
func PasswordHasher(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
