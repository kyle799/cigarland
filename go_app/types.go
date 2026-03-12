package main

import "time"

type Cigar struct {
	Brand                string `gorm:"primaryKey" json:"brand"`
	Name                 string `gorm:"primaryKey" json:"name"`
	Origin               string `json:"origin"`
	Wrapper              string `json:"wrapper"`
	Profile              string `json:"profile"`
	TastyTip             bool   `gorm:"default:false" json:"tasty_tip"`
	Pressed              bool   `gorm:"default:false" json:"pressed"`
	Binder               string `json:"binder"`
	Spicy                int    `gorm:"default:0" json:"spicy"`
	Rating               int    `json:"rating"`
	Length               int    `json:"length"`
	Ring                 int    `json:"ring"`
	Review               string `json:"review"`
	JohnRating           int    `json:"john_rating"`
	JohnReview           string `json:"john_review"`
	KyleRating           int    `json:"kyle_rating"`
	KyleReview           string `json:"kyle_review"`
	ImageRef             string `json:"image_ref"`
	AuthenticHumanReview string `json:"authentic_human_review"`
}

type CigarCreatePayload struct {
	Cigars []*Cigar `json:"cigar_list"`
}

type SelectionFilter struct {
	Column   string `json:"column" jsonschema:"required"`
	Value    any    `json:"value" jsonschema:"required"`
	Operator string `json:"operator" jsonschema:"required"`
	Logical  string `json:"logical"`
}

type QueryPayload struct {
	Table   string            `json:"table"`
	Filters []SelectionFilter `json:"filters"`
}

type Session struct {
	ID        string    `gorm:"primaryKey;column:id"`
	Email     string    `gorm:"not null;column:email"`
	CreatedAt time.Time `gorm:"not null;column:created_at"`
}

type UserPermission struct {
	Email     string `gorm:"primaryKey;column:email" json:"email"`
	CanAdd    bool   `gorm:"not null;default:false;column:can_add" json:"can_add"`
	CanEdit   bool   `gorm:"not null;default:false;column:can_edit" json:"can_edit"`
	CanDelete bool   `gorm:"not null;default:false;column:can_delete" json:"can_delete"`
	CanAdmin  bool   `gorm:"not null;default:false;column:can_admin" json:"can_admin"`
}
