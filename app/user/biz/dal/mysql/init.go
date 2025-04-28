package mysql

import (
	"fmt"
	"os"

	"zqzqsb.com/gomall/app/user/biz/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")

	if mysqlUser == "" || mysqlPassword == "" || mysqlHost == "" {
		panic("Missing required environment variables: MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlHost)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
}
