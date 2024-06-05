package model

import "gorm.io/datatypes"

type Book struct {
	ModelBase
	Title           string                      `json:"title"`
	Description     string                      `json:"description"`
	Author          string                      `json:"author"`
	PublicationYear int                         `json:"publication_year"`
	Genre           datatypes.JSONSlice[string] `json:"genre"`
	CoverImage      string                      `json:"cover_image"`
	Price           float64                     `json:"price"`
}
