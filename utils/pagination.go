package utils

import (
	"math"
	"strings"

	"github.com/jinzhu/gorm"
)

type Pagination struct {
	Page    int
	Limit   int
	OrderBy []string
	Preload []string
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

func (p *Pagination) Paginate(db *gorm.DB, m interface{}) *Resource {
	var meta MetaResource
	var offset int
	var count int

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 30
	}

	done := make(chan bool, 1)

	go func() {
		db.Model(m).Count(&count)
		done <- true
	}()

	order := strings.Join(p.OrderBy[:], ", ")
	order = strings.ReplaceAll(order, ".", " ")

	if len(p.Preload) > 0 {
		for _, include := range p.Preload {
			preload := ToCamelCase(include)
			db = db.Preload(preload)
		}
	}
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
