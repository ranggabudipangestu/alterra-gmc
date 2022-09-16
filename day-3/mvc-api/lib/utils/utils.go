package utils

import "github.com/ranggabudipangestu/mvc-api/models"

func RemoveIndex(s []models.Book, index int) []models.Book {
	return append(s[:index], s[index+1:]...)
}
