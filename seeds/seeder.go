package seeds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/localhearts/wms/models"
	"gorm.io/gorm"
)

func Load(db *gorm.DB, filename string) error {
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
		&models.PurchaseOrder{},
		&models.PurchaseOrderDetail{},
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

	// isi data uom melalu file json
	jsonFile, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("gagal membuka file JSON: %w", err)
	}
	defer jsonFile.Close()

	// Baca seluruh isi file
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("gagal membaca file JSON: %w", err)
	}

	// Unmarshal JSON ke slice Uom
	var uoms []models.Uom
	if err := json.Unmarshal(byteValue, &uoms); err != nil {
		return fmt.Errorf("gagal melakukan unmarshal data JSON: %w", err)
	}

	// Sisipkan data secara batch ke database
	if err := db.Create(&uoms).Error; err != nil {
		return fmt.Errorf("gagal menyisipkan data UOM: %w", err)
	}

	log.Println("✅ Seeding Uom completed successfully.")

	return nil
}
