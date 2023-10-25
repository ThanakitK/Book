package repositories

import "golang-db/models"

type BookRepository interface {
	GetAll() (result []models.RepoBookModel, err error)
	Create(payload models.RepoBookModel) (err error)
	Update(book_id string, payload models.UpdateRepoBookModel) (err error)
	Delete(book_id string) (err error)
}
