package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/denisK-H/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

const (
	queryNames       = "SELECT name FROM users"
	queryUniqueNames = "SELECT DISTINCT name FROM users"

	errQueryPrefix = "db query:"
	errScanPrefix  = "rows scanning:"
	errRowsPrefix  = "rows error:"
)

var (
	errDB   = errors.New("db error")
	errScan = errors.New("scan error")
)

func TestGetNames(t *testing.T) {
	t.Parallel()

	d, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer d.Close()

	service := db.New(d)

	mock.ExpectQuery(queryNames).
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow("Alice").
				AddRow("Bob"),
		)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Alice", "Bob"}, names)

	mock.ExpectQuery(queryNames).
		WillReturnError(errDB)

	_, err = service.GetNames()
	require.ErrorContains(t, err, errQueryPrefix)

	mock.ExpectQuery(queryNames).
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).AddRow(123),
		)

	_, err = service.GetNames()
	require.ErrorContains(t, err, errScanPrefix)

	mock.ExpectQuery(queryNames).
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow("Alice").
				RowError(0, errScan),
		)

	_, err = service.GetNames()
	require.ErrorContains(t, err, errScanPrefix)

	mock.ExpectQuery(queryNames).
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow("Alice").
				AddRow("Bob").
				CloseError(errScan),
		)

	_, err = service.GetNames()
	require.ErrorContains(t, err, errRowsPrefix)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	d, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer d.Close()

	service := db.New(d)

	mock.ExpectQuery(queryUniqueNames).
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow("Alice").
				AddRow("Bob"),
		)

	names, err := service.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Alice", "Bob"}, names)

	mock.ExpectQuery(queryUniqueNames).
		WillReturnError(errDB)

	_, err = service.GetUniqueNames()
	require.ErrorContains(t, err, errQueryPrefix)

	mock.ExpectQuery(queryUniqueNames).
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).AddRow(123),
		)

	_, err = service.GetUniqueNames()
	require.ErrorContains(t, err, errScanPrefix)

	mock.ExpectQuery(queryUniqueNames).
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow("Alice").
				RowError(0, errScan),
		)

	_, err = service.GetUniqueNames()
	require.ErrorContains(t, err, errScanPrefix)

	mock.ExpectQuery(queryUniqueNames).
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow("Alice").
				AddRow("Bob").
				CloseError(errScan),
		)

	_, err = service.GetUniqueNames()
	require.ErrorContains(t, err, errRowsPrefix)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_RowIterationError(t *testing.T) {
	t.Parallel()

	d, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer d.Close()

	service := db.New(d)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob").
		RowError(1, errDB)

	mock.ExpectQuery(queryNames).
		WillReturnRows(rows)

	_, err = service.GetNames()
	require.Error(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_RowIterationError(t *testing.T) {
	t.Parallel()

	d, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer d.Close()

	service := db.New(d)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob").
		RowError(1, errDB)

	mock.ExpectQuery(queryUniqueNames).
		WillReturnRows(rows)

	_, err = service.GetUniqueNames()
	require.Error(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_NilValue(t *testing.T) {
	t.Parallel()

	d, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer d.Close()

	service := db.New(d)

	mock.ExpectQuery(queryNames).
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).AddRow(nil),
		)

	names, err := service.GetNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, errScanPrefix)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_NilValue(t *testing.T) {
	t.Parallel()

	d, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer d.Close()

	service := db.New(d)

	mock.ExpectQuery(queryUniqueNames).
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).AddRow(nil),
		)

	names, err := service.GetUniqueNames()
	require.Nil(t, names)
	require.ErrorContains(t, err, errScanPrefix)

	require.NoError(t, mock.ExpectationsWereMet())
}
