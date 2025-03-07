package database

import (
	"fmt"

	"github.com/localhearts/wms/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "mysql" {
		dsn := DbUser + ":" + DbPassword + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
		server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database")
		} else {
			fmt.Println("Connected to database")
		}
	}
	server.DB.Debug().AutoMigrate(
		&models.Product{},
		&models.Supplier{},
		&models.Warehouse{},
		&models.Location{})
	server.DB.Debug().AutoMigrate(
		&models.Inbound{},
		&models.InboundDetail{},
	)
	server.DB.Debug().AutoMigrate(
		&models.Vas{},
		&models.Outbound{},
		&models.OutboundDetail{},
	)
	server.DB.Debug().AutoMigrate(
		&models.Stock{},
	)
}
