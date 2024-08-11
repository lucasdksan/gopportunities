package model

import "time"

type Opening struct {
	Id        uint      `json:"id"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
	CreatedAT time.Time `json:"createAt"`
	UpdatedAT time.Time `json:"updatedAt"`
	DeletedAT time.Time `json:"deletedAt,omitempty"`
}
