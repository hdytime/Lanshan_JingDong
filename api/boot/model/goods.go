package model

type Goods struct {
	ID       uint    `gorm:"primaryKey"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"quantity"`
	Reviews  uint    `json:"reviews"`
	GoodsID  int64   `json:"goods_id"`
	Content  string  `json:"content"`
}

type Reviews struct {
	ID       uint   `gorm:"primaryKey"`
	GoodsID  int64  `json:"goods_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

//type ReviewsForUser struct {
//	Username string `json:"username"`
//	Name     string `json:"name"`
//	Content  string `json:"content"`
//}
