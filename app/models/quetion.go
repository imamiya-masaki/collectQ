package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Quetion struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"context"`
}

func GetAllQuetion(db *gorm.DB) []Quetion {
	var quetions []Quetion
	db.Find(&quetions)
	fmt.Println("確認！！", quetions)
	return quetions
}

func GetQuetion(db *gorm.DB) Quetion {
	var quetion Quetion
	db.Find(&quetion)
	fmt.Println("確認!!T", quetion)
	return quetion
}

func (quetion *Quetion) CreateQuetion(db *gorm.DB) {
	result := db.Create(&quetion)
	if err := result.Error; err != nil {
		fmt.Println("err", err)
		// println(err)
	}
}
