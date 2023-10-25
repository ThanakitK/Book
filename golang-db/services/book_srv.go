package services

import (
	"errors"
	"golang-db/models"
	"golang-db/repositories"

	"github.com/google/uuid"
)

type bookSrv struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookService {
	return bookSrv{bookRepo}
}

func (s bookSrv) GetAllBook() (result []models.SrvBookModel, err error) {

	res, err := s.bookRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, v := range res {
		result = append(result, models.SrvBookModel(v))
	}

	return result, nil
}

func (s bookSrv) CreateBook(payload models.SrvBookModel) (err error) {

	if payload.Title == "" {
		return errors.New("title not found")
	}

	s.bookRepo.Create(models.RepoBookModel{
		Id:    uuid.New().String(),
		Title: payload.Title,
		Price: payload.Price,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s bookSrv) UpdateBook(book_id string, payload models.UpdateSrvBookModel) (err error) {
	s.bookRepo.Update(book_id, models.UpdateRepoBookModel{
		Title: payload.Title,
		Price: payload.Price,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s bookSrv) DeleteBook(book_id string) (err error) {
	s.bookRepo.Delete(book_id)

	if err != nil {
		return err
	}

	return nil
}
