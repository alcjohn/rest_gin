package models

type Book struct {
	BaseModel
	Title  string `gorm:"Not Null"json:"title"`
	Author string `gorm:"Not Null"json:"author"`
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func (b *Book) ValidateCreateInput() {
	if err := 
}
