package service

import (
	"challengeGO/model"
	"challengeGO/repository"
)

type CategoryService interface {
	Create(category *model.Category) error
	GetAll() ([]model.Category, error)
	Update(category *model.Category) error
	Delete(id uint) error
	FindByID(id uint) (*model.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{r}
}

func (s *categoryService) Create(category *model.Category) error {
	return s.repo.Create(category)
}

func (s *categoryService) GetAll() ([]model.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) Update(category *model.Category) error {
	return s.repo.Update(category)
}

func (s *categoryService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *categoryService) FindByID(id uint) (*model.Category, error) {
	return s.repo.FindByID(id)
}
