package store

import "apiserver/internal/model"

// MovieRepository - указание на хранилище.
type MovieRepository struct {
	store *Store
}

// Create - создание нового элемента в базе данных
func (r *MovieRepository) Create(m *model.Movie) (*model.Movie, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO movie (name_movie, year_movie, contry) VALUES ($1, $2, $3) RETURNING id",
		m.NameMovie,
		m.YearMovie,
		m.Contry,
	).Scan(&m.ID); err != nil {
		return nil, err
	}
	return m, nil
}

// FindByMovieName ...
func (r *MovieRepository) FindByMovieName(nameMovie string) (*model.Movie, error) {
	m := &model.Movie{}
	if err := r.store.db.QueryRow(
		"SELECT id, name_movie, year_movie, contry FROM movie WHERE name_movie=$1",
		nameMovie,
	).Scan(&m.ID,
		m.NameMovie,
		m.YearMovie,
		m.Contry,
	); err != nil {
		return nil, err
	}

	return m, nil
}
