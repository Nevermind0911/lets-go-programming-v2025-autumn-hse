package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	mockErrSelect = errors.New("select query failure")
	mockErrRows   = errors.New("row iteration failure")
)

const (
	userAlice = "Alice"
	userBob   = "Bob"
)

func TestGetNames_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		AddRow(userBob)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := s.GetNames()
	require.NoError(t, err)
	assert.Equal(t, []string{userAlice, userBob}, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_Empty(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"})

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := s.GetNames()
	require.NoError(t, err)
	assert.Empty(t, names)
}

func TestGetNames_QueryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	mock.ExpectQuery("SELECT name FROM users").WillReturnError(mockErrSelect)

	names, err := s.GetNames()
	assert.ErrorContains(t, err, "db query")
	assert.Nil(t, names)
}

func TestGetNames_ScanFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := s.GetNames()
	assert.ErrorContains(t, err, "rows scanning")
	assert.Nil(t, names)
}

func TestGetNames_RowsIterationFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).RowError(0, mockErrRows)

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := s.GetNames()
	assert.ErrorContains(t, err, "rows error")
	assert.Nil(t, names)
}

func TestGetUniqueNames_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		AddRow(userBob)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	assert.Equal(t, []string{userAlice, userBob}, names)
}

func TestGetUniqueNames_Empty(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"})

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	assert.Empty(t, names)
}

func TestGetUniqueNames_QueryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(mockErrSelect)

	names, err := s.GetUniqueNames()
	assert.ErrorContains(t, err, "db query")
	assert.Nil(t, names)
}

func TestGetUniqueNames_ScanFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	assert.ErrorContains(t, err, "rows scanning")
	assert.Nil(t, names)
}

func TestGetUniqueNames_RowsIterationFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).RowError(0, mockErrRows)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	assert.ErrorContains(t, err, "rows error")
	assert.Nil(t, names)
}
