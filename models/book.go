package models

type Book struct {
	BaseModel
	Title  string `gorm:"Not Null" json:"title"`
	Author string `gorm:"Not Null" json:"author"`
}
