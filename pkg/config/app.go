package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// This file returns a variable db, which helps interacting with the database
 
var (
	db *gorm.DB
)

func Connect(){
	// to connect with db
	d, err := gorm.Open("mysql", "shivi:password/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}