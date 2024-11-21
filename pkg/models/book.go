package models

import (
	"errors"
	"fmt"

	"github.com/Tahseen-Zaman/book-store-go/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		panic("Failed to migrate the database: " + err.Error())
	}
}

func (b *Book) CreateBook() *Book {
	err := db.Where(Book{Name: b.Name, Author: b.Author, Publication: b.Publication}).First(b).Error
	if err == nil {
		return b
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	err = db.Create(b).Error
	if err != nil {
		return  nil
	}
	return b
}

func GetAllBooks() ([]Book, error) {
    var books []Book
    err := db.Find(&books).Error
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve books: %w", err)
    }

    return books, nil
}


func GetBookById(Id int64) *Book {
	var book Book
	err := db.First(&book, Id).Error
	if err != nil {
		return nil
	}
	return &book
}

func DeleteBook(Id int64) error {
	err := db.Delete(&Book{}, Id).Error
	if err != nil {
		return err
	}
	return nil
}