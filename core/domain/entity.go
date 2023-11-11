package domain

import "time"

type Entity struct {
	ID        string    `json:"id" valid:"notnull" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at" valid:"-" gorm:"column:created_at;type:time"`
	UpdatedAt time.Time `json:"updated_at" valid:"-" gorm:"column:updated_at;type:time"`
}
