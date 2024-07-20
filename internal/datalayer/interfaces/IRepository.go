package interfaces

import (
	"golang.org/x/exp/constraints"
)

type idType interface {
	constraints.Integer | ~string
}

type IRepository[T any, Id idType] interface {
	Get(id Id) *T
	GetAll() []T
	Create(data T) *T
	Delete(id Id) *T
	Put(data T) *T
}
