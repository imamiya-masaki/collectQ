package models

import "gorm.io/gorm"

type Quetion struct {
	id         int    `json:"id"`
	title      string `json:"title"`
	content    string `json:"context"`
	created_at int    `json:"created_at"`
}

func GetAllQuetion(db *gorm.DB) []Quetion {
	var quetions []Quetion
	db.Find(&quetions)
	return quetions
}

func (quetion *Quetion) CreateQuetion(db *gorm.DB) {
	result := db.Create(&quetion)
	if err := result.Error; err != nil {
		println(err)
	}
}
