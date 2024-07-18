package interfaces

import "GoApi/internal/domain"

type IAlbumServices interface {
	Get(id int64) *domain.Album
	GetAll() []domain.Album
	Create(album domain.Album) *domain.Album
	Delete(id int64) *domain.Album
	Put(album domain.Album) *domain.Album
}
