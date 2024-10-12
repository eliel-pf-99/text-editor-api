package users

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func AddID(user User) User {
	user.ID = uuid.NewString()
	return user
}

func ParseToUser(userReq UserSignUp) User {
	user := User{
		Email:    userReq.Email,
		Name:     userReq.Name,
		Password: userReq.Password,
	}
	return user
}

func HashPassword(user User) (User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return User{}, err
	}
	user.Password = string(hashed)
	return user, nil
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
