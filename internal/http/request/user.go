package request

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const UserIDLen = 8

// nolint: tagliatelle
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u User) Validate() error {
	if err := validation.ValidateStruct(&u,
		validation.Field(&u.ID, validation.Required, validation.Length(UserIDLen, UserIDLen), is.Digit),
		validation.Field(&u.FirstName, validation.Required, is.UTFLetterNumeric),
		validation.Field(&u.LastName, validation.Required, is.UTFLetterNumeric),
	); err != nil {
		return fmt.Errorf("user validation failed %w", err)
	}

	return nil
}
