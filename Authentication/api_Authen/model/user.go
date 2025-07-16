package model

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"time"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func CreateUser(db *gorm.DB, user *User) error {
	// Hash password befor add in database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println("Create user successful")
	return nil
}

func LoginUser(db *gorm.DB, user *User) (string, error) {
	// get user from email
	SelectedUser := new(User)
	result := db.Where("email = ?", user.Email).First(SelectedUser)
	if result.Error != nil {
		return "", result.Error
	}
	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(SelectedUser.Password), []byte(user.Password))
	if err != nil {
		return "", result.Error
	}

	// pass = return jwt
	jwtSecretKey := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = SelectedUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", nil
	}
	return t, nil
}
