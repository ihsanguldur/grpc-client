package models

type Todo struct {
	Content string `json:"content"`
	UserID  uint   `json:"userID"`
	Status  bool   `json:"status" gorm:"default:false"`
}
