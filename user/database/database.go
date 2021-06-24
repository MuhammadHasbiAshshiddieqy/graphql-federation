package db

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Order *gorm.DB;
var Apra *gorm.DB;

// InitDB - connect to db
func InitDB() {
    var err error
    Order, err = gorm.Open("mysql", os.Getenv("DATA_ORDER"))
    if err != nil {
        fmt.Println(err)
        panic("failed to connect order database")
    }

    Order.LogMode(true)

    Apra, err = gorm.Open("mysql", os.Getenv("DATA_APRA"))
    if err != nil {
        fmt.Println(err)
        panic("failed to connect apra database")
    }

    Apra.LogMode(true)
}
