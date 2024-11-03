package repository

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (r *Repository[T]) DeleteById(db *gorm.DB, entity *T, id any) error {
	return db.Delete(entity, "id = ?", id).Error
}

func (r *Repository[T]) FindById(db *gorm.DB, entity *T, id any) error {
	return db.Where("id = ?", id).Take(entity).Error
}

func (r *Repository[T]) FindByIdPreloadAll(db *gorm.DB, entity *T, id any) error {
	return db.Where("id = ?", id).Preload(clause.Associations).Take(entity).Error
}

func (r *Repository[T]) GetAll(db *gorm.DB) ([]T, error) {
	var datas []T
	return datas, db.Model(new(T)).Find(&datas).Error
}

func (r *Repository[T]) FoundRecordById(db *gorm.DB, id any) (bool, error) {
	var count int64
	if err := db.Model(new(T)).Where("id = ?", id).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (r *Repository[T]) GetAllPagination(db *gorm.DB, p Pagination) ([]T, MetaPagination, error) {
	var metaPagin MetaPagination
	var datas []T

	queyChain := db.Model(datas)
	metaError := metaPagination(queyChain, p.Limit, &metaPagin)

	if metaError != nil {
		return datas, metaPagin, metaError
	}

	errFind := queyChain.Scopes(Paginate(p)).Find(&datas).Error
	if errFind != nil {
		return datas, metaPagin, errFind
	}

	return datas, metaPagin, nil
}

func (r *Repository[T]) PaginationAggregation(queryDB *gorm.DB, p Pagination) ([]T, MetaPagination, error) {
	var metaPagin MetaPagination
	var datas []T

	metaError := metaPagination(queryDB, p.Limit, &metaPagin)

	if metaError != nil {
		return datas, metaPagin, metaError
	}

	errFind := queryDB.Scopes(Paginate(p)).Find(&datas).Error
	if errFind != nil {
		return datas, metaPagin, errFind
	}

	return datas, metaPagin, nil
}
