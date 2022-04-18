package functions

import (
	model "authapp/packages/coponents/user/models"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

func ValidateUser(user model.User) []string {
	errs := []string{}
	isValidMail := regexp.MustCompile(emailRegex).MatchString(user.Email)
	if !isValidMail {
		errs = append(errs, "Invalid Email")
	}
	if len(user.Password) < 4 {
		errs = append(errs, "Invalid Password; Possword must be more than 4 characters")
	}
	if len(user.Name) < 1 {
		errs = append(errs, "Please enter a name")
	}
	return errs
}

func GestHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}