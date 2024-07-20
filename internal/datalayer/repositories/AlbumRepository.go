package repositories

import "GoApi/internal/domain"

type AlbumRepository struct{}

func NewAlbumRepository() *AlbumRepository {
	return &AlbumRepository{}
}

var albums = []domain.Album{
	{Id: 1, Price: 13.99, Artist: "Oleg F. A.", Name: "Her'"},
	{Id: 2, Price: 10.99, Artist: "Manis K. K.", Name: "YouNga"},
	{Id: 3, Price: 9.99, Artist: "Alister K. E.", Name: "Shorts"},
	{Id: 4, Price: 2.99, Artist: "Capitan K. E.", Name: "Manual"},
}

func (ar *AlbumRepository) Get(id int64) *domain.Album {
	for _, a := range albums {
		if a.Id == id {
			return &a
		}
	}
	return nil
}

func (ar *AlbumRepository) GetAll() []domain.Album {
	return albums
}

func (ar *AlbumRepository) Create(data domain.Album) *domain.Album {
	data.Id = int64(len(albums) + 1)
	albums = append(albums, data)
	return &data
}

func (ar *AlbumRepository) Delete(id int64) *domain.Album {
	for i, a := range albums {
		if a.Id == id {
			deletedAlbum := &albums[i]
			albums = append(albums[:i], albums[i+1:]...)
			return deletedAlbum
		}
	}
	return nil
}

func (ar *AlbumRepository) Put(data domain.Album) *domain.Album {
	for i, a := range albums {
		if a.Id == data.Id {
			albums[i] = data
			return &albums[i]
		}
	}
	albums = append(albums, data)
	return &data
}
