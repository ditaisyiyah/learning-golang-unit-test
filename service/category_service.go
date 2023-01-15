package service

import (
	"errors"
	entity "learning-golang-unit-test/entity"
	repository "learning-golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id) // remember! a value of interface type can hold any value that implements those methods
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
