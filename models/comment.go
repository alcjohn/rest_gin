package models

type Comment struct {
	BaseModel
	User    *User  `json:"user,omitempty"`
	UserID  uint   `json:"user_id,omitempty"`
	Content string `gorm:"Not Null" json:"content"`
	BookID  uint   `gorm:"Not Null" json:"book_id"`
}
