package model

import "time"

type User struct {
	// identity
	ID        string `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`

	// changes
	Endpoints []Endpoint `bson:"endpoints"`
}

type Endpoint struct {
	// only on registration
	URL       string        `bson:"url"`
	Interval  time.Duration `bson:"interval"`
	Threshold int           `bson:"threshold"`

	// continously updated
	Remain  int `bson:"failed"`
	Fail    int `bson:"fail"`
	Success int `bson:"success"`
}
