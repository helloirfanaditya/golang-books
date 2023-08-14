package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"micros/database"
	"micros/models"
)

func CreateBook(payload models.Book) (models.Book, error) {
	q := "INSERT INTO books (title,description,price) VALUES ($1,$2,$3)"

	if payload.Title == "" || payload.Description == "" || payload.Price == 0 {
		return models.Book{}, errors.New("Please fill the title / description / price")
	}

	_, err := database.Db.Exec(q, payload.Title, payload.Description, payload.Price)

	if err != nil {
		return models.Book{}, err
	}

	return payload, nil

}

func GetBooks() ([]models.Book, error) {
	var payload []models.Book
	q := "SELECT * FROM books"
	result, err := database.Db.Query(q)
	if err != nil {
		return nil, err
	}
	for result.Next() {
		var book models.Book
		err := result.Scan(&book.ID, &book.Title, &book.Description, &book.Price)
		if err != nil {
			return nil, err
		}
		payload = append(payload, book)
	}
	if len(payload) == 0 {
		return []models.Book{}, err
	}
	return payload, nil
}

func FindBook(payload int) (models.Book, error) {
	q := "SELECT * FROM books WHERE id = $1"
	row := database.Db.QueryRow(q, payload)

	var book models.Book
	err := row.Scan(&book.ID, &book.Title, &book.Description, &book.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, fmt.Errorf("No Row Found")
		}
		return models.Book{}, fmt.Errorf("error scanning row: %w", err)
	}

	return book, nil
}

func UpdateBook(payload models.Book) (models.Book, error) {
	q := "UPDATE books SET title=$1,description=$2,price=$3 WHERE id = $4"

	if payload.Title == "" || payload.Description == "" || payload.Price == 0 || payload.ID == 0 {
		return models.Book{}, errors.New("Please fill the title / description / price")
	}

	_, err := database.Db.Exec(q, payload.Title, payload.Description, payload.Price, payload.ID)

	if err != nil {
		return models.Book{}, errors.New("failed to update book: " + err.Error())
	}
	return payload, nil
}

func DeleteBook(payload models.Book) (models.Book, error) {
	row := "SELECT id FROM books WHERE id=$1"
	result := database.Db.QueryRow(row, payload.ID)
	errRow := result.Scan(&payload.ID)

	if errRow != nil {
		return models.Book{}, errors.New("Row Not Found")
	}

	q := "DELETE FROM books WHERE id=$1"
	_, errDelete := database.Db.Exec(q, payload.ID)

	if errDelete != nil {
		return models.Book{}, errors.New("failed to delete book: " + errDelete.Error())
	}
	return models.Book{}, nil
}
