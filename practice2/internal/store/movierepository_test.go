package store_test

import (
	"apiserver/internal/model"
	"apiserver/internal/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovieRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("movie")

	u, err := s.Movie().Create(&model.Movie{
		NameMovie: "Testing",
		YearMovie: "test",
		Contry:    "contryTest",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestMovieRepository_FindByName(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("movie")

	movieName := "Testing"
	_, err := s.Movie().FindByMovieName(movieName)
	assert.Error(t, err)

	s.Movie().Create(&model.Movie{
		NameMovie: "Testing",
		YearMovie: "test",
		Contry:    "contryTest",
	})

	m, err := s.Movie().FindByMovieName(movieName)
	assert.NoError(t, err)
	assert.NotNil(t, m)

}
