package seeds

import "gorm.io/gorm"

func Load(db *gorm.DB) {
	err := db.Debug().AutoMigrate(
		&Supplier{},
		&Product{},
	)
	if err != nil {
		panic(err)
	}
}
