package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string
	Password string
}

var users = map[string]string{}

func RegisterUser(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	users[email] = string(hashedPassword)
	log.Println(users)
	return nil
}
func AuthenticateUser(email, password string) bool {
	hashedPassword, ok := users[email]
	if !ok {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
