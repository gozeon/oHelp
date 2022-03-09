package model

type Order struct {
	CommonModel
	Title       string `gorm:"size:120;index" binding:"required" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Tags        string `gorm:"size:150;index" json:"tags"`
	Status      bool   `gorm:"index" json:"status"`
	CreateUser  string `gorm:"size:25;not null"  json:"create_user"`
}

type Comment struct {
	CommonModel
	Description string `gorm:"type:text" json:"description"`
	CreateUser  string `gorm:"size:25;not null"  json:"create_user"`
	OrderId     uint   `json:"order_id"`
}
