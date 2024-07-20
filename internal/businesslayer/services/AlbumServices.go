package services

import (
	"GoApi/internal/datalayer/interfaces"
	"GoApi/internal/domain"
)

type AlbumServices struct {
	albumRepository interfaces.IRepository[domain.Album, int64]
}

func NewAlbumServices(repository interfaces.IRepository[domain.Album, int64]) *AlbumServices {
	return &AlbumServices{
		albumRepository: repository,
	}
}

func (as *AlbumServices) Get(id int64) *domain.Album {
	return as.albumRepository.Get(id)
}

func (as *AlbumServices) GetAll() []domain.Album {
	return as.albumRepository.GetAll()
}

func (as *AlbumServices) Create(album domain.Album) *domain.Album {
	return as.albumRepository.Create(album)
}

func (as *AlbumServices) Put(album domain.Album) *domain.Album {
	return as.albumRepository.Put(album)
}

func (as *AlbumServices) Delete(id int64) *domain.Album {
	return as.albumRepository.Delete(id)
}
