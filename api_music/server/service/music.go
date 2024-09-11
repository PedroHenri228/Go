package service

import (
	"database/sql"
)

type Music struct {
	ID int
	Titulo string
	Artista string
	Album string
	Lancamento string
	Genero string
}

type MusicService struct {
	db *sql.DB
}


func NewMusicService(db *sql.DB) *MusicService {
	return &MusicService{db: db}
}

func(s *MusicService) SetMusic(music *Music) error {
	query := "INSERT INTO music(titulo, artista, album, lancamento, genero) VALUES (?,?,?,?,?)"

	result, err := s.db.Exec(query, music.Titulo, music.Artista, music.Album, music.Lancamento, music.Genero)

	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return err
	}

	music.ID = int(lastInsertID)

	return nil
}

func (s *MusicService) GetMusic() ([]Music, error) {
	query := "SELECT * FROM music"

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	var musics []Music

	for rows.Next() {
		var music Music
		err := rows.Scan(&music.ID, &music.Titulo, &music.Artista, &music.Album, &music.Lancamento, &music.Genero)

		if err != nil {
			return nil, err
		}

		musics = append(musics, music)
	}

	return musics, nil

	
}

func (s *MusicService) GetMusicByID(id int) (*Music, error) {
	query := "SELECT * FROM music WHERE id = ?"
	row :=  s.db.QueryRow(query, id)

	var music Music

	err := row.Scan(&music.ID, &music.Titulo, &music.Artista, &music.Album, &music.Lancamento, &music.Genero)

	if err != nil {
		return nil, err
	}

	return &music, nil

}

func (s *MusicService) UpdateMusic(music *Music) error {
	query := "UPDATE music SET titulo=?, artista=?, album=?, lancamento=?, genero=? WHERE id = ?"

	_, err := s.db.Exec(query, music.Titulo, music.Artista, music.Album, music.Lancamento, music.Genero, music.ID)


	return err
}

func (s *MusicService) DeleteMusic(id int) error {
	query := "DELETE FROM music WHERE id=?"
	_, err := s.db.Exec(query, id)
	return err
}