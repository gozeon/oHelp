package model

type Approval struct {
	CommonModel
	Title      string `gorm:"size:120;index" binding:"required" json:"title"`
	Content    string `gorm:"type:text"  binding:"required" json:"content"`
	Status     uint   `gorm:"index" json:"status" desc:"审评中 1、通过 2、驳回 3"`
	Approver   string `gorm:"size:25;index;not null"  binding:"required" json:"approver"`
	Assigner   string `gorm:"size:200" json:"assigner"`
	CreateUser string `gorm:"size:25;not null"  json:"create_user"`
}

type Progress struct {
	CommonModel
	Operate    uint   `gorm:"index" json:"operate" desc:"申请 1、转移 2、通过 3、驳回 4、回复 5"`
	Opinion    string `gorm:"type:text" json:"opinion"`
	CreateUser string `gorm:"size:25;not null"  json:"create_user"`
	ApprovalId uint   `json:"approval_id" binding:"required"`
}
