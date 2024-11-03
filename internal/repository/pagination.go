package repository

import (
	"math"
	"strings"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit   int
	Page    int
	Sort    string
	OrderBy string
}

type MetaPagination struct {
	TotalPages int
	TotalRows  int
}

func Paginate(p Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := p.Page
		if page <= 0 {
			page = 1
		}

		limit := p.Limit
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		pOrder := strings.ToUpper(p.Sort)
		order := "ASC"
		if pOrder == "ASC" || pOrder == "DESC" {
			order = pOrder
		}

		if p.OrderBy == "" {
			p.OrderBy = "created_at"
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit).Order(p.OrderBy + " " + order)
	}
}

func metaPagination(dbLogic *gorm.DB, limit int, meta *MetaPagination) error {
	var tRows int64

	err := dbLogic.Count(&tRows).Error

	meta.TotalRows = int(tRows)
	totalPages := int(math.Ceil(float64(tRows) / float64(limit)))
	meta.TotalPages = totalPages

	return err
}
