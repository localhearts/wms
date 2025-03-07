
type Warehouse struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null;unique"`
	Location  string    `gorm:"not null"`
	Area      float64   `gorm:"not null"`
	Capacity  int       `gorm:"not null"`
	Manager   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Storages  []Storage
}

type Storage struct {
	ID          uint      `gorm:"primaryKey"`
	WarehouseID uint      `gorm:"index"`
	Location    string    `gorm:"not null"`
	Rack        string    `gorm:"not null"`
	Level       int       `gorm:"not null"`
	Bin         int       `gorm:"not null"`
	Capacity    int       `gorm:"not null"`
	CurrentLoad int       `gorm:"not null"`
	Temperature float64
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type CategoryProduct struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Product struct {
	ID              uint      `gorm:"primaryKey"`
	Name            string    `gorm:"not null;unique"`
	CategoryID      uint      `gorm:"index"`
	SKU             string    `gorm:"unique;not null"`
	Description     string
	UnitOfMeasure   string    `gorm:"not null"`
	Weight          float64   `gorm:"not null"`
	Dimensions      string    `gorm:"not null"`
	ExpiryDate      time.Time
	ReorderLevel    int       `gorm:"not null"`
	StockMovements  []StockMovement
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
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