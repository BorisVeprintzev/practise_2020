package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config          *Config
	db              *sql.DB
	movieRepository *MovieRepository
}

// New ..
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

//Open - инициализация хранилища
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

//Close - действия при завершеннии работы
func (s *Store) Close() {
	s.db.Close()
}

//Movie - функция для работы с другими пакетами
func (s *Store) Movie() *MovieRepository {
	if s.movieRepository != nil {
		return s.movieRepository
	}

	s.movieRepository = &MovieRepository{
		store: s,
	}

	return s.movieRepository
}

// store.Movie().Create() - пример работы из вне
