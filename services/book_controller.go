package myapi

import (
	"errors"
	"log"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var (
	ErrInvalidBookId   = errors.New("invalid book id")
	ErrBookDoesntExist = errors.New("This book doesnt exist")
)

func GetBookById(params martini.Params, log *log.Logger, r render.Render) {
	idParam := params["id"]

	var (
		bookId int
		err    error
		book   *Book
	)

	if bookId, err = strconv.Atoi(idParam); err == nil {
		if book = FindBookById(bookId); book == nil {
			log.Println("Book does not exist")
			err = ErrBookDoesntExist
		}

	} else {
		err = ErrInvalidBookId
	}

	responseCode := 200

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
		responseCode = 500
	}

	r.JSON(responseCode, map[string]interface{}{
		"message": book,
		"error":   errMsg,
	})
}

func CreateBook() []byte {
	return nil
}
