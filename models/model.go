package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Response struct {
	Message []string    `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var db *gorm.DB
var err error

func init(){

	db, err = gorm.Open("postgres", "postgres://lxwpeswjghmixi:9ec7528ac10224d072ec3034b7c572890bbe321ca949c64c64b1440332207d22@ec2-54-83-48-188.compute-1.amazonaws.com:5432/dcg4m4p3grtu8b")
	db.SingularTable(true)
	if err != nil {
		log.Panic(err)
	}

}