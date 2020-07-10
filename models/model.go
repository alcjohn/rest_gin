package models

import (
	"math"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type Params struct {
	Page    int
	Limit   int
	OrderBy []string
}

type MetaResource struct {
	Page  int `json:"page"`
	Last  int `json:"last"`
	Total int `json:"total"`
	Limit int `json:"limit"`
}

type Resource struct {
	Data interface{}  `json:"data"`
	Meta MetaResource `json:"meta"`
}

func Paginate(db *gorm.DB, p *Params, m interface{}) *Resource {
	var meta MetaResource

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 30
	}

	order := strings.Join(p.OrderBy[:], ", ")
	order = strings.ReplaceAll(order, ".", " ")

	done := make(chan bool, 1)
	var count int
	var offset int

	go func() {
		db.Model(m).Count(&count)
		done <- true
	}()

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}
	db.Order(order).Limit(p.Limit).Offset(offset).Find(m)
	<-done

	last := int(math.Ceil(float64(count) / float64(p.Limit)))
	if last > 0 {
		meta.Last = last
	} else {
		meta.Last = 1
	}

	meta.Limit = p.Limit
	meta.Total = count
	meta.Page = p.Page

	return &Resource{
		Meta: meta,
		Data: m,
	}

}
