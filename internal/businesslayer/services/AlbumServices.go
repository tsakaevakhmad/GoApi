package services

import (
	"GoApi/internal/domain"
)

var albums = []domain.Album{
	{Id: 1, Price: 13.99, Artist: "Oleg F. A.", Name: "Her'"},
	{Id: 2, Price: 10.99, Artist: "Manis K. K.", Name: "YouNga"},
	{Id: 3, Price: 9.99, Artist: "Alister K. E.", Name: "Shorts"},
	{Id: 4, Price: 2.99, Artist: "Capitan K. E.", Name: "Manual"},
}

type AlbumServices struct{}

func NewAlbumServices() *AlbumServices {
	return &AlbumServices{}
}

func (l *AlbumServices) Get(id int64) *domain.Album {
	for _, a := range albums {
		if a.Id == id {
			return &a
		}
	}
	return nil
}

func (l *AlbumServices) GetAll() []domain.Album {
	return albums
}

func (l *AlbumServices) Create(album domain.Album) *domain.Album {
	album.Id = int64(len(albums) + 1)
	albums = append(albums, album)
	return &album
}

func (l *AlbumServices) Put(album domain.Album) *domain.Album {
	for _, a := range albums {
		if a.Id == album.Id {
			a = album
			return &a
		}
	}
	albums = append(albums, album)
	return &album
}

func (l *AlbumServices) Delete(id int64) *domain.Album {
	for i, a := range albums {
		if a.Id == id {
			deleteAlbum := &albums[i]
			albums = append(albums[:i], albums[i+1:]...)
			return deleteAlbum
		}
	}
	return nil
}
