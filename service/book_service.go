package service

import (
	"challengeGO/model"
	"challengeGO/repository"
)

type BookService interface {
	Create(book *model.Book) error
	GetByUser(userID uint) ([]model.Book, error)
	FindByID(id uint) (*model.Book, error) // tambahkan ini
	Update(book *model.Book) error
	Delete(id uint) error
	GetAll() ([]model.Book, error)
}

type bookService struct {
	repo repository.BookRepository
}

func (s *bookService) FindByID(id uint) (*model.Book, error) {
	return s.repo.FindByID(id)
}

func NewBookService(r repository.BookRepository) BookService {
	return &bookService{r}
}

func (s *bookService) GetAll() ([]model.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) Create(book *model.Book) error {
	return s.repo.Create(book)
}

func (s *bookService) GetByUser(userID uint) ([]model.Book, error) {
	return s.repo.FindAllByUser(userID)
}

func (s *bookService) Update(book *model.Book) error {
	return s.repo.Update(book)
}

func (s *bookService) Delete(id uint) error {
	return s.repo.Delete(id)
}
