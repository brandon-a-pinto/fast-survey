package validation

import (
	"fmt"
	"regexp"

	"github.com/brandon-a-pinto/fast-survey/internal/schemas"
)

const (
	minNameLength     = 6
	minPasswordLength = 8
)

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func paramIsRequired(p, t string) error {
	return fmt.Errorf("param %s (%s) is required", p, t)
}

func CreateUserValidation(p *schemas.CreateUserParams) error {
	if p.Email == "" && p.Name == "" && p.Password == "" && p.PasswordConfirmation == "" {
		return fmt.Errorf("request body empty or invalid")
	}
	if p.Email == "" {
		return paramIsRequired("email", "string")
	}
	if p.Name == "" {
		return paramIsRequired("name", "string")
	}
	if p.Password == "" {
		return paramIsRequired("password", "string")
	}
	if p.PasswordConfirmation == "" {
		return paramIsRequired("passwordConfirmation", "string")
	}
	if p.Password != p.PasswordConfirmation {
		return fmt.Errorf("password does not match passwordConfirmation")
	}
	if !isEmailValid(p.Email) {
		return fmt.Errorf("email is invalid")
	}
	if len(p.Name) <= minNameLength {
		return fmt.Errorf("name must be at least %d characters", minNameLength)
	}
	if len(p.Password) <= minPasswordLength {
		return fmt.Errorf("password must be at least %d characters", minPasswordLength)
	}

	return nil
}
