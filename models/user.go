package models

import (
	token "Go_REST_API_wit_JWT_Authentication_in_Gin/utils/token"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/html"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255; not null; unique" json:"username"`
	Password string `gorm:"size:255; not null"   json:"password"`
}

func GetUserByID(uid uint) (User, error) {
	var user User
	if err := DB.First(&user, uid).Error; err != nil {
		return User{}, errors.New("User not found!")
	}
	user.PrepareGive()
	return user, nil
}

func (user *User) PrepareGive() {
	user.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {
	var err error
	user := User{}
	err = DB.Model(User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token_res, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token_res, nil
}

func (user *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave() error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// remove spaces in username
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}
