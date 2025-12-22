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
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	expectedRows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		AddRow(userBob)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(expectedRows)

	result, err := s.GetNames()

	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, []string{userAlice, userBob}, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_Empty(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	emptyRows := sqlmock.NewRows([]string{"name"})

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(emptyRows)

	result, err := s.GetNames()

	require.NoError(t, err)
	assert.Empty(t, result)
}

func TestGetNames_QueryFail(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(mockErrSelect)

	list, err := s.GetNames()

	assert.ErrorContains(t, err, "db query")
	assert.Nil(t, list)
}

func TestGetNames_ScanFail(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rowsWithNull := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rowsWithNull)

	_, err = s.GetNames()
	assert.ErrorContains(t, err, "rows scanning")
}

func TestGetNames_RowsIterationFail(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	brokenRows := sqlmock.NewRows([]string{"name"}).
		AddRow(userAlice).
		RowError(1, mockErrRows)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(brokenRows)

	_, err = s.GetNames()
	assert.ErrorContains(t, err, "rows error")
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

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	result, err := s.GetUniqueNames()
	require.NoError(t, err)
	assert.Equal(t, []string{userAlice, userBob}, result)
}

func TestGetUniqueNames_Empty(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	emptyRows := sqlmock.NewRows([]string{"name"})

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(emptyRows)

	result, err := s.GetUniqueNames()
	require.NoError(t, err)
	assert.Empty(t, result)
}

func TestGetUniqueNames_QueryFail(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(mockErrSelect)

	_, err = s.GetUniqueNames()
	assert.ErrorContains(t, err, "db query")
}

func TestGetUniqueNames_ScanFail(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	_, err = s.GetUniqueNames()
	assert.ErrorContains(t, err, "rows scanning")
}

func TestGetUniqueNames_RowsIterationFail(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := New(db)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("SingleUser").
		RowError(1, mockErrRows)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	_, err = s.GetUniqueNames()
	assert.ErrorContains(t, err, "rows error")
}
