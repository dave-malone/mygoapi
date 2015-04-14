package myapi

import (
	"fmt"
	"log"
)

func FindBookById(bookId int) (book *Book) {
	log.Println("Getting book with id", bookId)

	db := GetDb()

	var (
		id           int    = *new(int)
		name, author string = *new(string), *new(string)
	)

	if err := db.QueryRow("SELECT id, name, author FROM book WHERE id = ?", bookId).Scan(&id, &name, &author); err == nil {
		book = new(Book)
		book.Id = id
		book.Name = name
		book.Author = author

		fmt.Printf("Found Book: %v\n", book)

	}
	return
}
