




type CategoryProduct struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}



type CycleCount struct {
	ID          uint      `gorm:"primaryKey"`
	WarehouseID uint      `gorm:"index"`
	StorageID   uint      `gorm:"index"`
	ProductID   uint      `gorm:"index"`
	CountedQty  int       `gorm:"not null"`
	ExpectedQty int       `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	Status      string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type Supplier struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null;unique"`
	Contact   string
	Email     string    `gorm:"unique"`
	Address   string
	City      string
	Country   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Inbound struct {
	ID             uint      `gorm:"primaryKey"`
	SupplierID     uint      `gorm:"index"`
	WarehouseID    uint      `gorm:"index"`
	Date           time.Time `gorm:"not null"`
	Status         string    `gorm:"not null"`
	ReferenceNo    string    `gorm:"not null;unique"`
	InboundDetails []InboundDetail
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

type InboundDetail struct {
	ID         uint      `gorm:"primaryKey"`
	InboundID  uint      `gorm:"index"`
	ProductID  uint      `gorm:"index"`
	Quantity   int       `gorm:"not null"`
	StorageID  uint      `gorm:"index"`
	ReceivedAt time.Time `gorm:"autoCreateTime"`
}

type Outbound struct {
	ID              uint      `gorm:"primaryKey"`
	WarehouseID     uint      `gorm:"index"`
	Date            time.Time `gorm:"not null"`
	Status          string    `gorm:"not null"`
	ReferenceNo     string    `gorm:"not null;unique"`
	OutboundDetails []OutboundDetail
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}

type OutboundDetail struct {
	ID         uint      `gorm:"primaryKey"`
	OutboundID uint      `gorm:"index"`
	ProductID  uint      `gorm:"index"`
	Quantity   int       `gorm:"not null"`
	StorageID  uint      `gorm:"index"`
	ShippedAt  time.Time `gorm:"autoCreateTime"`
}

type StockMovement struct {
	ID           uint      `gorm:"primaryKey"`
	ProductID    uint      `gorm:"index"`
	WarehouseID  uint      `gorm:"index"`
	StorageID    uint      `gorm:"index"`
	Quantity     int       `gorm:"not null"`
	MovementType string    `gorm:"not null"`
	Date         time.Time `gorm:"not null"`
	ReferenceNo  string    `gorm:"not null;unique"`
}

type Shipment struct {
	ID         uint      `gorm:"primaryKey"`
	RouteID    uint      `gorm:"index"`
	Status     string    `gorm:"not null"`
	TrackingNo string    `gorm:"unique"`
	Carrier    string    `gorm:"not null"`
	EstimatedDelivery time.Time
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type RoutePlanning struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null;unique"`
	Distance  float64
	Duration  int       `gorm:"not null"`
	Stops     int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type WavePicking struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Status    string    `gorm:"not null"`
	BatchSize int       `gorm:"not null"`
}

type DeliveryOrder struct {
	ID          uint      `gorm:"primaryKey"`
	OrderDate   time.Time `gorm:"not null"`
	Status      string    `gorm:"not null"`
	Customer    string    `gorm:"not null"`
	TotalAmount float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}   

Route Planning (Perencanaan Rute)
Optimasi rute pengiriman dilakukan berdasarkan faktor-faktor berikut:
✅ Jarak & waktu tempuh tercepat
✅ Kondisi lalu lintas & jalan (misal: jalan kecil, larangan truk)
✅ Beban kendaraan & efisiensi bahan bakar
✅ Jumlah drop point (multi-stop deliveries)

Sistem dapat menggunakan metode seperti:
📍 Shortest Path Algorithm → Menentukan rute tercepat berdasarkan jarak.
📍 Vehicle Routing Problem (VRP) → Mengoptimalkan pengiriman multi-stop.
📍 Time Window Routing → Menyesuaikan dengan jadwal penerimaan pelanggan.