package main

import (
//	"github.com/mattn/go-sqlite3"
//	"gorm.io/driver/sqlite"
//
// "gorm.io/gorm"
)

type Cigar struct {
	Brand                string `gorm:"primaryKey" json:"brand"`
	Name                 string `gorm:"primaryKey" json:"name"`
	Wrapper              string `json:"wrapper"`
	Profile              string `json:"profile"`
	TastyTip             bool   `gorm:"default:false" json:"tasty_tip"`
	Pressed              bool   `grom:"default:false" json:"pressed"`
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
