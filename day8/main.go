package main

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

type User struct {
	Name     string
	Email    string
	Password string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func ValidateUser(user User) error {
	if user.Name == "" {
		err := &ValidationError{
			Field:   "Name",
			Message: "Name cannot be empty",
		}

		return err
	}

	return nil
}

func main() {

	user := User{
		Name:     "",
		Email:    "user@example.com",
		Password: "password123",
	}

	err := ValidateUser(user)
	if err != nil {
		fmt.Println("Validation Error:", err)
	}

}
