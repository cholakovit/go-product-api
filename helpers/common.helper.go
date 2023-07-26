package helpers

import (
	"fmt"
	"log"
	"products/queries"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

// place is not here
func VerifyPassword(userPassword string, providedPassword string)(bool, string) {

	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("password is incorrect")
		check = false
	}
	return check, msg
}

func IsValidEmail(email *string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(emailPattern, *email)
	if err != nil {
		return false
	}
	return match
}

func VerifyEmail(email *string) string {

	isValidEmail := IsValidEmail(email)
	if isValidEmail == false {
		return "the email is not valid."
	}

	count, err := queries.FindUserByEmailQuery(email)
	if err != nil {
		return "error while searching for the user."
	}

	if count > 0 {
		return "this email is taken"
	}

	return ""
}