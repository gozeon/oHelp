package model

type Approval struct {
	CommonModel
	Title      string `gorm:"size:120;index" binding:"required" json:"title"`
	Content    string `gorm:"type:text" json:"content"`
	Status     string `gorm:"size:150;index" json:"tags" desc:"申请 1、审评中 2、通过 3、驳回 4"`
	Approver   string `gorm:"size:25;index;not null"  json:"approver"`
	Assigner   string `gorm:"size:200" json:"assigner"`
	CreateUser string `gorm:"size:25;not null"  json:"create_user"`
}

type Progress struct {
	CommonModel
	Operate    string `gorm:"size:150;index" json:"operate" desc:"转移 2、通过 3、驳回 4"`
	Opinion    string `gorm:"type:text" json:"content"`
	CreateUser string `gorm:"size:25;not null"  json:"create_user"`
	ApprovalId uint   `json:"approval_id"`
}
