package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Nevermind0911/task-6/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	errSelect = errors.New("select query failure")
	errRows   = errors.New("row iteration failure")
	errClose  = errors.New("rows close failure")
)

const (
	userAlice = "Alice"
	userBob   = "Bob"

	querySelectNames       = "SELECT name FROM users"
	querySelectUniqueNames = "SELECT DISTINCT name FROM users"
)

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

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

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"})

	mock.ExpectQuery(querySelectNames).WillReturnRows(rows)

	names, err := s.GetNames()
	require.NoError(t, err)
	assert.Empty(t, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_QueryFail(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

	mock.ExpectQuery(querySelectNames).WillReturnError(errSelect)

	names, err := s.GetNames()
	require.ErrorContains(t, err, "db query")
	assert.Nil(t, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_ScanFail(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery(querySelectNames).WillReturnRows(rows)

	names, err := s.GetNames()
	require.ErrorContains(t, err, "rows scanning")
	assert.Nil(t, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_RowsIterationFail(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		RowError(1, errRows)

	mock.ExpectQuery(querySelectNames).WillReturnRows(rows)

	names, err := s.GetNames()
	require.ErrorContains(t, err, "rows error")
	assert.Nil(t, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_CloseFail(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		CloseError(errClose)

	mock.ExpectQuery(querySelectNames).WillReturnRows(rows)

	names, err := s.GetNames()
	require.NoError(t, err)
	assert.Equal(t, []string{userAlice}, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_Success(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		AddRow(userBob)

	mock.ExpectQuery(querySelectUniqueNames).WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	assert.Equal(t, []string{userAlice, userBob}, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_Empty(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"})

	mock.ExpectQuery(querySelectUniqueNames).WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	assert.Empty(t, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_QueryFail(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

	mock.ExpectQuery(querySelectUniqueNames).WillReturnError(errSelect)

	names, err := s.GetUniqueNames()
	require.ErrorContains(t, err, "db query")
	assert.Nil(t, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_ScanFail(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery(querySelectUniqueNames).WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	require.ErrorContains(t, err, "rows scanning")
	assert.Nil(t, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_RowsIterationFail(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		RowError(1, errRows)

	mock.ExpectQuery(querySelectUniqueNames).WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	require.ErrorContains(t, err, "rows error")
	assert.Nil(t, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_CloseFail(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	s := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		CloseError(errClose)

	mock.ExpectQuery(querySelectUniqueNames).WillReturnRows(rows)

	names, err := s.GetUniqueNames()
	require.NoError(t, err)
	assert.Equal(t, []string{userAlice}, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}
