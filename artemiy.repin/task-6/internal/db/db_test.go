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

	querySelectNames       = "SELECT name FROM users"
	querySelectUniqueNames = "SELECT DISTINCT name FROM users"
)

func TestGetNames_Success(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		AddRow(userBob)

	mock.ExpectQuery(querySelectNames).WillReturnRows(rows)

	names, err := s.GetNames()
	require.NoError(t, err)
	assert.Equal(t, []string{userAlice, userBob}, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_Empty(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"})

	mock.ExpectQuery(querySelectNames).WillReturnRows(rows)

	names, err := s.GetNames()
	require.NoError(t, err)
	assert.Empty(t, names)
}

func TestGetNames_QueryFail(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	mock.ExpectQuery(querySelectNames).WillReturnError(mockErrSelect)

	names, err := s.GetNames()
	assert.ErrorContains(t, err, "db query")
	assert.Nil(t, names)
}

func TestGetNames_ScanFail(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery(querySelectNames).WillReturnRows(rows)

	names, err := s.GetNames()
	assert.ErrorContains(t, err, "rows scanning")
	assert.Nil(t, names)
}

func TestGetNames_RowsIterationFail(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		RowError(1, mockErrRows)

	mock.ExpectQuery(querySelectNames).WillReturnRows(rows)

	names, err := s.GetNames()
	assert.ErrorContains(t, err, "rows error")
	assert.Nil(t, names)
}

func TestGetUniqueNames_Success(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		AddRow(userBob)

	mock.ExpectQuery(querySelectUniqueNames).WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	assert.Equal(t, []string{userAlice, userBob}, names)
}

func TestGetUniqueNames_Empty(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"})

	mock.ExpectQuery(querySelectUniqueNames).WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	assert.Empty(t, names)
}

func TestGetUniqueNames_QueryFail(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	mock.ExpectQuery(querySelectUniqueNames).WillReturnError(mockErrSelect)

	names, err := s.GetUniqueNames()
	assert.ErrorContains(t, err, "db query")
	assert.Nil(t, names)
}

func TestGetUniqueNames_ScanFail(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery(querySelectUniqueNames).WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	assert.ErrorContains(t, err, "rows scanning")
	assert.Nil(t, names)
}

func TestGetUniqueNames_RowsIterationFail(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		RowError(1, mockErrRows)

	mock.ExpectQuery(querySelectUniqueNames).WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	assert.ErrorContains(t, err, "rows error")
	assert.Nil(t, names)
}
