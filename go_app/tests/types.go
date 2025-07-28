package main

type Cigar struct {
	Brand                string `gorm:"primaryKey" json:"brand" jsonschema:"required"`
	Name                 string `gorm:"primaryKey" json:"name" jsonschema:"required"`
	Wrapper              string `json:"wrapper,omitempty"`
	Profile              string `json:"profile,omitempty"`
	TastyTip             bool   `gorm:"default:false" json:"tasty_tip,omitempty"`
	Pressed              bool   `grom:"default:false" json:"pressed,omitempty"`
	Binder               string `json:"binder"`
	Spicy                int    `gorm:"default:0" json:"spicy,omitempty"`
	Rating               int    `json:"rating,omitempty"`
	Length               int    `json:"length,omitempty"`
	Ring                 int    `json:"ring,omitempty"`
	Review               string `json:"review,omitempty"`
	JohnRating           int    `json:"john_rating,omitempty"`
	JohnReview           string `json:"john_review,omitempty"`
	KyleRating           int    `json:"kyle_rating,omitempty"`
	KyleReview           string `json:"kyle_review,omitempty"`
	ImageRef             string `json:"image_ref,omitempty"`
	AuthenticHumanReview string `json:"authentic_human_review,omitempty"`
}

type CigarCreatePayload struct {
	Cigars []*Cigar `json:"cigar_list" jsonschema:"required"`
}
