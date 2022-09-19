package static

import (
	"errors"

	"github.com/ranggabudipangestu/mvc-api/lib/utils"
	"github.com/ranggabudipangestu/mvc-api/models"
)

var books []models.Book

func GetBooks() ([]models.Book, error) {
	if len(books) == 0 {
		return nil, errors.New("data not found")
	}
	return books, nil

}

func CreateBook(book models.Book) error {
	books = append(books, book)
	return nil

}

func GetBookById(id int) (*models.Book, error) {
	var bookData models.Book
	for _, data := range books {
		if data.ID == id {
			bookData.ID = data.ID
			bookData.Title = data.Title
			bookData.Description = data.Title
			bookData.Author = data.Author
		}
	}
	return &bookData, nil
}

func UpdateBook(id int, book models.Book) error {
	bookData, _ := GetBookById(id)
	DeleteBook(id)
	newBookData := models.Book{
		ID:          bookData.ID,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
	}

	books = append(books, newBookData)
	return nil

}

func DeleteBook(id int) error {
	var index int
	for i, data := range books {
		if data.ID == id {
			index = i
			break
		}
	}
	books = utils.RemoveIndex(books, index)

	return nil
}
