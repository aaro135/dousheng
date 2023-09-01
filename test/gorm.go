package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGorm(t *testing.T) {
	dsn := "root:Qwe1234...@tcp(127.0.0.1:3306)/simpletiktok?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Fatal(err)
	}

}
