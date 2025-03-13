package seeds

import (
	"log"

	"github.com/localhearts/wms/models"
	"gorm.io/gorm"
)

func Load(db *gorm.DB) {
	// Drop tables (order doesn't matter much when dropping)
	tablesMigrate := []interface{}{
		&models.Uom{},
		&models.Product{},
		&models.Category{},
		&models.City{},
		&models.Province{},
		&models.Supplier{},
		&models.Customer{},
		&models.Warehouse{},
		&models.Storage{},
		&models.Stock{},
		&models.PurchaseOrder{},
		&models.PurchaseOrderDetail{},
		&models.Inbound{},
		&models.InboundDetail{},
		&models.Vas{},
		&models.DeliveryOrder{},
		&models.DeliveryOrderDetail{},
	}

	for _, table := range tablesMigrate {
		if db.Migrator().HasTable(table) {
			if err := db.Debug().Migrator().DropTable(table); err != nil {
				panic(err)
			}
			log.Printf("Dropped table for model: %T", table)
		}
	}

	// Migrate tables in proper dependency order.
	// Note that Supplier must be created before Product (and any other models referencing it).
	err := db.Debug().AutoMigrate(tablesMigrate...)

	if err != nil {
		panic(err)
	}
	log.Println("✅ Migration completed successfully.")
}
