package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/sam3016/bookshelfapi/internal/database"
)

func (apiCfg *apiConfig) handlerCreateBook(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Title       string    `json:"title"`
		Volume      int32     `json:"volume"`
		Category    string    `json:"category"`
		Author      string    `json:"author"`
		PublishedAt time.Time `json:"published_at"`
		Publisher   string    `json:"publisher"`
		Finished    bool      `json:"finished"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	book, err := apiCfg.DB.CreateBook(r.Context(), database.CreateBookParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Title:       params.Title,
		Volume:      params.Volume,
		Category:    params.Category,
		Author:      params.Author,
		PublishedAt: params.PublishedAt,
		Publisher:   params.Publisher,
		Finished:    params.Finished,
		UserID:      user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create book: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseBookToBook(book))
}

func (apiCfg *apiConfig) handlerGetBooks(w http.ResponseWriter, r *http.Request, user database.User) {
	books, err := apiCfg.DB.GetBooks(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't get books: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseBooksToBooks(books))
}

func (apiCfg *apiConfig) handlerDeleteBook(w http.ResponseWriter, r *http.Request, user database.User) {
	bookIDStr := chi.URLParam(r, "bookID")
	bookID, err := uuid.Parse(bookIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't delete book id: %v", err))
		return
	}

	err = apiCfg.DB.DeleteBook(r.Context(), database.DeleteBookParams{
		ID:     bookID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't delete book: %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}

func (apiCfg *apiConfig) handlerUpdateBook(w http.ResponseWriter, r *http.Request, user database.User) {
	bookIDStr := chi.URLParam(r, "bookID")
	bookID, err := uuid.Parse(bookIDStr)

	type parameters struct {
		Title       string    `json:"title"`
		Volume      int32     `json:"volume"`
		Category    string    `json:"category"`
		Author      string    `json:"author"`
		PublishedAt time.Time `json:"published_at"`
		Publisher   string    `json:"publisher"`
		Finished    bool      `json:"finished"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err1 := decoder.Decode(&params)
	if err1 != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	err2 := apiCfg.DB.UpdateBook(r.Context(), database.UpdateBookParams{
		ID:          bookID,
		Title:       params.Title,
		Volume:      params.Volume,
		Category:    params.Category,
		Author:      params.Author,
		PublishedAt: params.PublishedAt,
		Publisher:   params.Publisher,
		Finished:    params.Finished,
		UserID:      user.ID,
	})
	if err2 != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't update book: %v", err))
		return
	}

	// respondWithJSON(w, 200, databaseBookToBook(book))
}
