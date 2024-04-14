package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/sam3016/bookshelfapi/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Book struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Volume      int32     `json:"volume"`
	Category    string    `json:"category"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Publisher   string    `json:"publisher"`
	Finished    bool      `json:"finished"`
	UserID      uuid.UUID `json:"user_id"`
}

func databaseBookToBook(dbBook database.Book) Book {
	return Book{
		ID:          dbBook.ID,
		CreatedAt:   dbBook.CreatedAt,
		UpdatedAt:   dbBook.UpdatedAt,
		Title:       dbBook.Title,
		Volume:      dbBook.Volume,
		Category:    dbBook.Category,
		Author:      dbBook.Author,
		PublishedAt: dbBook.PublishedAt,
		Publisher:   dbBook.Publisher,
		Finished:    dbBook.Finished,
		UserID:      dbBook.UserID,
	}
}

func databaseBooksToBooks(dbBooks []database.Book) []Book {
	books := []Book{}
	for _, dbBook := range dbBooks {
		books = append(books, databaseBookToBook(dbBook))
	}
	return books
}
