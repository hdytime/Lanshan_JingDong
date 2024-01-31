package model

type Goods struct {
	ID       uint    `gorm:"primaryKey"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
	Reviews  string  `json:"reviews"`
}
