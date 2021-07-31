package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func validatePassword(hashedPassword []byte, password string) {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Password check passed!")
	}
}

func main() {
	password := `Iam(omple*Pa$$w0rD`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("\nHashedPassword is", string(hashedPassword))

	fmt.Println("\nChecking if given password matches hash, case success")
	validatePassword(hashedPassword, password)

	fmt.Println("\nChecking if given password matches hash, case failure ")
	wrongPassword := `I am wrong pass`
	validatePassword(hashedPassword, wrongPassword)
}
