package models

import (
	"math"
	"time"
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

func Paginate(p *Params, m interface{}) *Resource {
	var meta MetaResource

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 30
	}

	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			DB.Order(o)
		}
	}

	done := make(chan bool, 1)
	var count int
	var offset int

	go func() {
		DB.Model(m).Count(&count)
		done <- true
	}()

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}
	DB.Limit(p.Limit).Offset(offset).Find(m)
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