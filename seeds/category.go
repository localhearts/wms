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

// SeedUom seeder untuk data Uom.
func SeedCategory(db *gorm.DB, filename string) error {
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

	// Unmarshal JSON ke slice Category
	var categories []models.Category
	if err := json.Unmarshal(byteValue, &categories); err != nil {
		return fmt.Errorf("gagal melakukan unmarshal data JSON: %w", err)
	}

	// Sisipkan data secara batch ke database
	if err := db.Create(&categories).Error; err != nil {
		return fmt.Errorf("gagal menyisipkan data Category Product: %w", err)
	}

	log.Println("✅ Seeding Category Product completed successfully.")

	return nil
}
