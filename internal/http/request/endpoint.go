package request

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Endpoint struct {
	URL       string        `bson:"url"`
	Interval  time.Duration `bson:"interval"`
	Threshold int           `bson:"threshold"`
}

func (ep Endpoint) Validate() error {
	if err := validation.ValidateStruct(&ep,
		validation.Field(&ep.URL, validation.Required, is.URL),
		// validation.Field(&ep.Interval, validation.Required, is.Duration),
		// validation.Field(&ep.Threshold, validation.Required, is.Digit),
	); err != nil {
		return fmt.Errorf("endpoint validation failed %w", err)
	}

	return nil
}
