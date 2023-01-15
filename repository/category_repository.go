package repository

import (
	entity "learning-golang-unit-test/entity"
)

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
