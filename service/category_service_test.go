package service

import (
	"learning-golang-unit-test/entity"
	"learning-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	categoryRepository = &repository.CategoryRepositoryMock{}
	categoryService    = CategoryService{Repository: categoryRepository}
)

func TestCategoryService_GetNotFound(t *testing.T) {

	/// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	result, err := categoryService.Get("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Juned",
	}

	/// program mock
	categoryRepository.Mock.On("FindById", "1").Return(category)

	result, err := categoryService.Get("1")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
