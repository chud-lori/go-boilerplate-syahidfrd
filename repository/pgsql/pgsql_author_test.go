package pgsql_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/syahidfrd/go-boilerplate/entity"
	"github.com/syahidfrd/go-boilerplate/repository/pgsql"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestAuthorRepo_Create(t *testing.T) {
	author := &entity.Author{
		Name:      "name",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "INSERT INTO authors"
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(author.Name, author.CreatedAt, author.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	authorRepo := pgsql.NewPgsqlAuthorRepository(db)
	err = authorRepo.Create(context.TODO(), author)
	assert.NoError(t, err)
}

func TestAuthorRepo_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	authorMock := entity.Author{
		ID:        1,
		Name:      "name",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(authorMock.ID, authorMock.Name, authorMock.CreatedAt, authorMock.UpdatedAt)

	query := "SELECT id, name, created_at, updated_at FROM authors WHERE id = $1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(1).
		WillReturnRows(rows)

	authorRepo := pgsql.NewPgsqlAuthorRepository(db)
	author, err := authorRepo.GetByID(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, author)
	assert.Equal(t, authorMock.ID, author.ID)
}

func TestAuthorRepo_Fetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockAuthors := []entity.Author{
		{ID: 1, Name: "name", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "name 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(mockAuthors[0].ID, mockAuthors[0].Name, mockAuthors[0].CreatedAt, mockAuthors[0].UpdatedAt).
		AddRow(mockAuthors[1].ID, mockAuthors[1].Name, mockAuthors[1].CreatedAt, mockAuthors[1].UpdatedAt)

	query := "SELECT id, name, created_at, updated_at FROM authors"
	mock.ExpectQuery(query).WillReturnRows(rows)

	authorRepo := pgsql.NewPgsqlAuthorRepository(db)
	authors, err := authorRepo.Fetch(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, authors, 2)
}

func TestAuthorRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	author := &entity.Author{
		ID:        1,
		Name:      "name",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := "UPDATE authors SET name = $1, updated_at = $2 WHERE id = $3"
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(author.Name, author.UpdatedAt, author.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	authorRepo := pgsql.NewPgsqlAuthorRepository(db)
	err = authorRepo.Update(context.TODO(), author)
	assert.NoError(t, err)
}

func TestAuthorRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "DELETE FROM authors WHERE id = $1"
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	authorRepo := pgsql.NewPgsqlAuthorRepository(db)
	err = authorRepo.Delete(context.TODO(), 1)
	assert.NoError(t, err)
}
