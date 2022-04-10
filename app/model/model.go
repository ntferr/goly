package model

import (
	"fmt"

	"github.com/ntferr/goly/app/internal/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Goly struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Goly     string `json:"goly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		settings.GetSettings().Database.Host, settings.GetSettings().Database.User,
		settings.GetSettings().Database.Pwd, settings.GetSettings().Database.Name, settings.GetSettings().Database.Port)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Goly{})
	if err != nil {
		fmt.Println(err)
	}
}
