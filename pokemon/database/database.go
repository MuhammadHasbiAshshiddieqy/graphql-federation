package db

import (
  "log"
  "time"
	"fmt"
	"os"

    // Import go-sql-driver
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  "gorm.io/gorm/logger"
)

// Pokemon variable for gorm.DB database pokemon
var Pokemon *gorm.DB;

// InitDB - connect to db
func InitDB() {
    var err error
    Pokemon, err = gorm.Open(mysql.Open(os.Getenv("DATA_POKEMON")), &gorm.Config{
        Logger: loggerSetting(),
        PrepareStmt: true,
    })
    if err != nil {
        fmt.Println(err)
        panic("failed to connect pokemon database")
    }
}

func loggerSetting() (logger.Interface) {
    newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
            SlowThreshold:              time.Second,    // Slow SQL threshold
            LogLevel:                   logger.Info,    // Log level
            IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
            Colorful:                  true,            // Disable color
		},
	)

    return newLogger
}
