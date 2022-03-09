package model

type Comment struct {
	CommonModel
	Description string `gorm:"type:text" json:"description"`
	CreateUser  string `gorm:"size:25;not null"  json:"create_user"`
	OrderId     uint   `json:"order_id"`
}
