package services

import "golang-db/models"

type BookService interface {
	GetAllBook() (result []models.SrvBookModel, err error)
	CreateBook(payload models.SrvBookModel) (err error)
	UpdateBook(book_id string, payload models.UpdateSrvBookModel) (err error)
	DeleteBook(book_id string) (err error)
}
