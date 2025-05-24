package repository

import (
	"challengeGO/model"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *model.Book) error
	FindAllByUser(userID uint) ([]model.Book, error)
	Update(book *model.Book) error
	Delete(id uint) error
	FindByID(id uint) (*model.Book, error)
	FindAll() ([]model.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) Create(book *model.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) FindAllByUser(userID uint) ([]model.Book, error) {
	var books []model.Book
	err := r.db.Preload("Category").Where("user_id = ?", userID).Find(&books).Error
	return books, err
}

func (r *bookRepository) Update(book *model.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := r.db.Preload("Category").Find(&books).Error
	return books, err
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&model.Book{}, id).Error
}

func (r *bookRepository) FindByID(id uint) (*model.Book, error) {
	var book model.Book
	err := r.db.Preload("Category").First(&book, id).Error // preload Category supaya terisi
	return &book, err
}
