package seeds

import (
	"log"

	"github.com/localhearts/wms/models"
	"gorm.io/gorm"
)

func Load(db *gorm.DB) {
	err := db.Debug().Migrator().DropTable(&models.Product{}, &models.Supplier{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Product{}, &models.Supplier{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
