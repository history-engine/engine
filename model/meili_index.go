package model

import "time"

// MeiliIndex 索引信息
type MeiliIndex struct {
	Uid        string    `json:"uid,omitempty"`
	PrimaryKey string    `json:"primaryKey,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}
