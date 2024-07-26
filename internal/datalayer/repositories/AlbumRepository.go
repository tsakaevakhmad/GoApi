package repositories

import (
	"GoApi/internal/datalayer"
	"GoApi/internal/domain"
	"database/sql"
	"log"
)

type AlbumRepository struct {
	context *sql.DB
}

func NewAlbumRepository(driver string, connection string) *AlbumRepository {
	return &AlbumRepository{
		context: datalayer.GetContext(driver, connection),
	}
}

var albums = []domain.Album{
	{Id: 1, Price: 13.99, Artist: "Oleg F. A.", Name: "Her'"},
	{Id: 2, Price: 10.99, Artist: "Manis K. K.", Name: "YouNga"},
	{Id: 3, Price: 9.99, Artist: "Alister K. E.", Name: "Shorts"},
	{Id: 4, Price: 2.99, Artist: "Capitan K. E.", Name: "Manual"},
}

func (ar *AlbumRepository) Get(id int64) *domain.Album {
	var album domain.Album
	query := `SELECT * FROM albums WHERE id = $1`
	ar.context.QueryRow(query, id).Scan(&album.Id, &album.Price, &album.Artist, &album.Name)
	return &album
}

func (ar *AlbumRepository) GetAll() []domain.Album {
	var albums []domain.Album
	query := `SELECT * FROM albums`
	rows, err := ar.context.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var album domain.Album
		err := rows.Scan(&album.Id, &album.Price, &album.Artist, &album.Name)

		if err != nil {
			log.Fatal(err)
			return nil
		}

		albums = append(albums, album)
	}
	return albums
}

func (ar *AlbumRepository) Create(data domain.Album) *domain.Album {
	query := `INSERT INTO Albums(price, artist, name) VALUES($1, $2, $3) RETURNING id`
	err := ar.context.QueryRow(query, data.Price, data.Artist, data.Name).Scan(&data.Id)
	if err != nil {
		log.Fatal(err)
	}
	return &data
}

func (ar *AlbumRepository) Delete(id int64) *domain.Album {
	var album domain.Album
	query := `DELETE FROM albums WHERE id = $1 RETURNING *`
	ar.context.QueryRow(query, id).Scan(&album.Id, &album.Price, &album.Artist, &album.Name)
	return &album
}

func (ar *AlbumRepository) Put(data domain.Album) *domain.Album {
	query := `UPDATE albums SET price=$1, artist=$2, name=$3 WHERE id=$4 RETURNING *`
	ar.context.QueryRow(query, data.Price, data.Artist, data.Name, data.Id).Scan(&data.Id, &data.Price, &data.Artist, &data.Name)
	return &data
}
