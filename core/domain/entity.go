package domain

import "time"

type Entity struct {
	ID        string    `json:"id" valid:"notnull"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}
