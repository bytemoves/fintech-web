package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword (password string) ( string , error) {
	hash , err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	if err != nil{
		return "",err
	}

	return string(hash),nil
}


func VerifyPassword (password,hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password),[]byte(password))
	return err
}